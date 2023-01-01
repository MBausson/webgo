package server

import (
	"log"
	"net"
	"webgo/pkg/Http"
	"webgo/pkg/Resources"
)

func SetHost(hostt, portt string) {
	host = hostt
	port = portt
}

func AddEndpoint(ep string, ru Resources.WebResource) {
	rru := Resources.Endpoint{
		EndpointRoute: ep,
		Wr:            &ru,
		Method:        Http.GET,
	}

	endpoints = append(endpoints, &rru)
}

// AddDirectory
// loads every file in a given directory
// Directory : path to said directory
// Recursive : indicate whether every children directories should be loaded, or simply the first children of given directory
//
//	Endpoints strings are transformed as '/<filename.ext>' or '/<previous_directories>/<filename.ext> if recursive is set to true
func AddDirectory(dir string, recursive bool) {
	//	Grab the desired file names (relative to 'dir' directory)
	filenames := load_dirs(dir, "", recursive)

	//	Create a new webresource for each file name and add an endpoint to it
	for _, fname := range filenames {
		wr := Resources.WebResource{LocalPath: dir + "/" + fname}
		AddEndpoint("/"+fname, wr)
	}
}

func AddFuncEndpoint(ep string, f func(Http.HttpRequest) Http.HttpResponse, meth Http.HttpMethod) {
	fr := Resources.FuncResource{f, meth}

	rep := Resources.Endpoint{
		EndpointRoute: ep,
		Method:        meth,
		Fr:            &fr,
	}

	endpoints = append(endpoints, &rep)
}

func Start() {
	//	Before starting up the socket server, loads each resource unit
	load_resources()

	log.Println("*-*-*-* Creating server socket *-*-*-*")
	server, err := net.Listen("tcp", get_host_port())

	check_socket_error(err)
	defer server.Close()

	log.Printf("*-*-*-* Listening on %v *-*-*-*\n", server.Addr())

	for {
		connection, err := server.Accept()

		check_socket_error(err)
		go handle_client(connection)
	}
}
