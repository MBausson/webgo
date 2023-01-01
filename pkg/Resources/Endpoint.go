package Resources

import "webgo/pkg/Http"

/*	An endpoint is a resource at a specific route */
/*	The resource can be a FuncResource (defined by the developer, and called on request) */
/*	It can also be a WebResource (a file)	*/

type Endpoint struct {
	EndpointRoute string
	Method        Http.HttpMethod
	Wr            *WebResource
	Fr            *FuncResource
}
