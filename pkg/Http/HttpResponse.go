package Http

import (
	"encoding/json"
	"strings"
)

type HttpResponse struct {
	Version string
	Code    HttpCode
	Headers []HttpHeader
	Content string
}

func (res HttpResponse) ToString() string {
	result := res.Version + " " + strings.ToUpper(string(res.Code)) + "\r\n"

	for i, header := range res.Headers {
		result += header.ToString()

		if i != len(res.Headers)-1 {
			result += "\r\n"
		}
	}

	if len(res.Content) != 0 {
		result += "\r\n\r\n" + res.Content
	}

	return result
}

func DefaultOKtext(text string) HttpResponse {
	return HttpResponse{
		Version: "HTTP/1.1",
		Code:    OK,
		Headers: []HttpHeader{
			{"Content-Type", "text/plain"},
		},
		Content: text,
	}
}

func DefaultOKjstring(content string) HttpResponse {
	return HttpResponse{
		Version: "HTTP/1.1",
		Code:    OK,
		Headers: []HttpHeader{
			{"Content-Type", "text/json"},
		},
		Content: content,
	}
}

func DefaultOkjobj(obj any) HttpResponse {
	res, err := json.Marshal(obj)

	if err != nil {
		panic(err)
	}

	return DefaultOKjstring(string(res))
}
