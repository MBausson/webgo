package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"runtime"
	"strings"
	"webgo/pkg/Http"
	"webgo/pkg/Mime"
	"webgo/pkg/Resources"
)

var (
	host      string
	port      string
	endpoints = []*Resources.Endpoint{}
)

func get_host_port() string {
	return host + ":" + port
}

func check_socket_error(err error) {
	if err != nil {
		log.Panicf("\t!!! Error with socket : %v", err)
	}
}

func handle_client(client net.Conn) {
	defer close_connection(&client)
	log.Printf("> Received connection: [%v]\n", client.RemoteAddr())

	buffer := make([]byte, 1024)
	client.Read(buffer)
	//	Remove trailing NULL characters
	content := strings.Trim(string(buffer), string(0))

	//	If request is empty, end here
	if len(content) == 0 {
		return
	}

	//	Get the request object
	req, err := Http.RequestFromString(content)

	//	if the request could not be translated to request object, send error
	if err != nil {
		send_response(generic_unreadable(), client)
		return
	}

	//	Grab the desired endpoint
	asked, found := get_endpoint(req.EndpointRoute, req.Method)
	var response Http.HttpResponse

	log.Printf(">>> Received request from [%v]\n\t<%v> at '%v'", client.RemoteAddr(), req.Method, req.CompleteRoute)

	//	If the desired endpoint does not exist
	if found == Http.BadRequest {
		response = generic_not_found()
	} else if found == Http.MethodNotAllowed {
		//	If the request method does not correspond with the resource method
		response = generic_method_not_allowed(asked.Method, req.Method)
	} else {
		//	Here, everything is fine

		//	If the desired resource is a WebResource (a file)
		//	Just send the response with file's content as body
		if asked.Wr != nil {
			response = Http.HttpResponse{
				Version: "HTTP/1.1",
				Code:    Http.OK,
				Headers: []Http.HttpHeader{
					{Key: "Content-Type", Value: asked.Wr.ContentType},
					{Key: "Host", Value: get_host_port()},
				},
				Content: asked.Wr.GetResult(),
			}
		} else {
			//	Here, the desired resource is a user-func
			response = asked.Fr.GetResult(req)
		}
	}

	//	Sends the final response
	send_response(response, client)
}

func get_endpoint(route string, method Http.HttpMethod) (Resources.Endpoint, Http.HttpCode) {
	for _, ep := range endpoints {
		if ep.EndpointRoute == route {
			if ep.Method == method {
				return *ep, Http.OK
			} else {
				return *ep, Http.MethodNotAllowed
			}
		}
	}

	return Resources.Endpoint{}, Http.BadRequest
}

func send_response(response Http.HttpResponse, conn net.Conn) {
	_, err := conn.Write([]byte(response.ToString()))
	check_socket_error(err)

	log.Printf("<<< Sent response to [%v]\n\t%v", conn.RemoteAddr(), response.Code)
}

func generic_not_found() Http.HttpResponse {
	return Http.HttpResponse{
		Version: "HTTP/1.1",
		Code:    Http.NotFound,
		Headers: []Http.HttpHeader{{"Content-Type", "text/html"}},
		Content: "Could not find requested resource.",
	}
}

func generic_method_not_allowed(expected, got Http.HttpMethod) Http.HttpResponse {
	return Http.HttpResponse{
		Version: "HTTP/1.1",
		Code:    Http.MethodNotAllowed,
		Headers: []Http.HttpHeader{{"Content-Type", "text/html"}},
		Content: fmt.Sprintf("Method not allowed : expected %v, got %v", expected, got),
	}
}

func generic_unreadable() Http.HttpResponse {
	return Http.HttpResponse{
		Version: "HTTP/1.1",
		Code:    Http.BadRequest,
		Headers: []Http.HttpHeader{
			{Key: "Content-Type", Value: Mime.GetDefault()},
			{Key: "Host", Value: get_host_port()},
		},
		Content: "Request could not be readt as a valid HTTP request.",
	}
}

func close_connection(connection *net.Conn) {
	log.Printf("< Closing connection [%v]", (*connection).RemoteAddr())
	err := (*connection).Close()
	check_socket_error(err)
}

func load_dirs(ogdir, curdir string, recursive bool) []string {
	entries, err := os.ReadDir(ogdir + curdir)

	if err != nil {
		log.Panicf(" !!! Could not open directory %v / or read the files in it.", ogdir+curdir)
	}

	result := []string{}

	for _, e := range entries {
		//	If the current entry is a directory
		//	... and if we don't want recursivity, skip this entry
		if e.IsDir() {
			if !recursive {
				continue
			}

			//	Otherwise, make a recursive call, and add to curdir the current directory name
			result = append(result, load_dirs(ogdir, curdir+"/"+e.Name(), true)...)
		} else {

			//	Here, the entry is a file
			var name string

			//	We decide the file name we want
			//	curdir == "" means we are in the 1st iteration of the function
			//	So the file name is just its name
			//	Otherwise, the filename is the current directory name + the file name
			//	note: with this method, the filename has a leading '/' character, explaining the [1:] slice
			if len(curdir) == 0 {
				name = e.Name()
			} else {
				name = (curdir + "/" + e.Name())[1:]
			}

			result = append(result, name)
		}
	}

	return result
}

func load_resources() {
	for _, ep := range endpoints {
		if (*ep).Wr != nil {
			(*ep).Wr.Load()

			log.Printf("--- ✔️ Loaded resource at '%v' [%v] ---\n", (*ep).Wr.LocalPath, ep.EndpointRoute)
		} else {
			log.Printf("--- ✔️ Loaded function '%v' [%v] ---\n", runtime.FuncForPC(reflect.ValueOf((*ep).Fr.F).Pointer()).Name(), ep.EndpointRoute)
		}
	}
}
