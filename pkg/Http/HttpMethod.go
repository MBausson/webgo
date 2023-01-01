package Http

import "strings"

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

func MethodFromString(content string) HttpMethod {
	content = strings.ToLower(content)

	switch content {
	case "get":
		return GET

	case "post":
		return POST

	case "put":
		return PUT

	case "patch":
		return PATCH

	case "delete":
		return DELETE
	}

	return ""
}
