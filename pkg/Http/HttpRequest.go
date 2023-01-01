package Http

import (
	"strings"
)

type HttpRequest struct {
	Version       string             //	HTTP version 'HTTP/X.Y'
	EndpointRoute string             //	Requested resource without URL parameters
	CompleteRoute string             //	Full URL request
	Method        HttpMethod         //	HTTP method (GET, PUT...)
	UrlParameters ParameterContainer //	Url parameters
	Headers       []HttpHeader       //	HTTP headers
	Content       string             //	Message content
}

func RequestFromString(content string) (HttpRequest, error) {
	lines := strings.Split(content, "\r\n")

	if len(lines) == 0 {
		return HttpRequest{}, &UnreadableRequest{}
	}

	//	1st line <Method Route Version>
	first_line := strings.Split(lines[0], " ")

	if len(first_line) != 3 {
		return HttpRequest{}, &UnreadableRequest{}
	}

	method := MethodFromString(first_line[0])
	route := first_line[1]
	version := first_line[2]

	//	endpoint url, without URL parameters
	endpoint_splitted := strings.Split(route, "?")
	endpoint_route := endpoint_splitted[0]
	endpoint_params := UrlToParameters(strings.Join(endpoint_splitted[1:], "?"))

	//	Headers
	headers := []HttpHeader{}

	if len(lines) == 1 {
		return HttpRequest{
			Method:        method,
			EndpointRoute: endpoint_route,
			CompleteRoute: route,
			UrlParameters: endpoint_params,
			Version:       version,
			Headers:       []HttpHeader{},
			Content:       "",
		}, nil
	}

	iline := 1

	for lines[iline] != "" {
		headers = append(headers, HeaderFromString(lines[iline]))
		iline++
	}

	//	Request body
	body := ""

	for j := iline + 1; j < len(lines); j++ {
		body += lines[j]

		//	Check if it's the last iteration, if it is, we would not want to add a line break
		if j != len(lines)-1 {
			body += "\r\n"
		}
	}

	return HttpRequest{
		Method:        method,
		EndpointRoute: endpoint_route,
		CompleteRoute: route,
		UrlParameters: endpoint_params,
		Version:       version,
		Headers:       headers,
		Content:       body,
	}, nil
}

func (req1 HttpRequest) Equals(req2 HttpRequest) bool {
	//	We can't use the default '==' operator because of the []HttpHeaders field
	if len(req1.Headers) != len(req2.Headers) {
		return false
	}

	for i := 0; i < len(req1.Headers); i++ {
		if req1.Headers[i] != req2.Headers[i] {
			return false
		}
	}

	return req1.Version == req2.Version &&
		req1.CompleteRoute == req2.CompleteRoute &&
		req1.Method == req2.Method &&
		req1.Content == req2.Content
}
