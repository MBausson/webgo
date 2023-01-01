package Resources

import "webgo/pkg/Http"

type FuncResource struct {
	F      func(Http.HttpRequest) Http.HttpResponse
	Method Http.HttpMethod
}

func (fr FuncResource) GetResult(req Http.HttpRequest) Http.HttpResponse {
	return fr.F(req)
}
