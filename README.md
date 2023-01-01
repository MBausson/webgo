#   Webgo

Webgo is a learning Go project, it is not intended to be used on production !  
Webgo is simple in its form : you can set up a web server with a few lines, specify web resources in a directory, and even handle HTTP requests via custom Go function  

##  Tutorial

First of all, the needed package is `webgo/internal/server`  
In order to set up your server, simply call `server.SetHost(<host>, <port>)`
The server is started up using the `server.Start()` function, which is blocking.  

### Specify reachable resources

Webgo lets the user defines directories of resources, which will be available from HTTP get requests  

```go
/*
    Directory name must be a path relative to current directory './'
    Note the 'true' argument which refers to the 'recursive' parameter, indicating that we want nested directories inside './my_dir' to be loaded.
*/
server.AddDirectory("./my_dir", true)

/*
    This function adds an endpoint to your server, from an individual file.
    The first argument is the URL to be reached for this resource, it should always start with a slash '/'
    The second parameter is the file resource (from webgo/pkg/Resources), we can only indicate the file path
*/
server.AddEndpoint("/homepage", Resources.WebResource{
	LocalPath: "./path/page.html",
})

/*
    AddFuncEndpoint lets you handle endpoint with a Go function
    The first parameter is the route, with a leading slash '/'
    The second is a function, taking in parameter a Http.HttpRequest and returning an Http.HttpResponse
    The last parameter, the HTTP method.
 */
server.AddFuncEndpoint("/echo", func(request Http.HttpRequest) Http.HttpResponse {
    return Http.DefaultOKtext(request.Content)
}, Http.GET)
```