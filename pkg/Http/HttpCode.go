package Http

type HttpCode string

const (
	Continue            HttpCode = "100 Continue"
	SwitchingProtocols  HttpCode = "101 Switching Protocols"
	OK                  HttpCode = "200 OK"
	Created             HttpCode = "201 Created"
	Accepted            HttpCode = "202 Accepted"
	NonAuthoritative    HttpCode = "203 Non-Authoritative Information"
	NoContent           HttpCode = "204 No Content"
	ResetContent        HttpCode = "205 Reset Content"
	PartialContent      HttpCode = "206 Partial Content"
	MultipleChoices     HttpCode = "300 Multiple Choices"
	MovedPermanently    HttpCode = "301 Moved Permanently"
	Found               HttpCode = "302 Found"
	SeeOther            HttpCode = "303 See Other"
	NotModified         HttpCode = "304 Not Modified"
	UseProxy            HttpCode = "305 Use Proxy"
	TemporaryRedirect   HttpCode = "307 Temporary Redirect"
	BadRequest          HttpCode = "400 Bad Request"
	Unauthorized        HttpCode = "401 Unauthorized"
	PaymentRequired     HttpCode = "402 Payment Required"
	Forbidden           HttpCode = "403 Forbidden"
	NotFound            HttpCode = "404 Not Found"
	MethodNotAllowed    HttpCode = "405 Method Not Allowed"
	NotAcceptable       HttpCode = "406 Not Acceptable"
	ProxyAuthRequired   HttpCode = "407 Proxy Authentication Required"
	RequestTimeout      HttpCode = "408 Request Timeout"
	Conflict            HttpCode = "409 Conflict"
	Gone                HttpCode = "410 Gone"
	LengthRequired      HttpCode = "411 Length Required"
	PreconditionFailed  HttpCode = "412 Precondition Failed"
	RequestEntityToo    HttpCode = "413 Request Entity Too Large"
	URITooLong          HttpCode = "414 Request-URI Too Long"
	UnsupportedMedia    HttpCode = "415 Unsupported Media Type"
	RequestedRange      HttpCode = "416 Requested Range Not Satisfiable"
	ExpectationFailed   HttpCode = "417 Expectation Failed"
	InternalServerError HttpCode = "500 Internal Server Error"
	NotImplemented      HttpCode = "501 Not Implemented"
	BadGateway          HttpCode = "502 Bad Gateway"
	ServiceUnavailable  HttpCode = "503 Service Unavailable"
	GatewayTimeout      HttpCode = "504 Gateway Timeout"
	HTTPVersionNot      HttpCode = "505 HTTP Version Not Supported"
)

func CodeFromString(code string) HttpCode {
	switch code {
	case "100":
		return Continue
	case "101":
		return SwitchingProtocols
	case "200":
		return OK
	case "201":
		return Created
	case "202":
		return Accepted
	case "203":
		return NonAuthoritative
	case "204":
		return NoContent
	case "205":
		return ResetContent
	case "206":
		return PartialContent
	case "300":
		return MultipleChoices
	case "301":
		return MovedPermanently
	case "302":
		return Found
	case "303":
		return SeeOther
	case "304":
		return NotModified
	case "305":
		return UseProxy
	case "307":
		return TemporaryRedirect
	case "400":
		return BadRequest
	case "401":
		return Unauthorized
	case "402":
		return PaymentRequired
	case "403":
		return Forbidden
	case "404":
		return NotFound
	case "405":
		return MethodNotAllowed
	case "406":
		return NotAcceptable
	case "407":
		return ProxyAuthRequired
	case "408":
		return RequestTimeout
	case "409":
		return Conflict
	case "410":
		return Gone
	case "411":
		return LengthRequired
	case "412":
		return PreconditionFailed
	case "413":
		return RequestEntityToo
	case "414":
		return URITooLong
	case "415":
		return UnsupportedMedia
	case "416":
		return RequestedRange
	case "417":
		return ExpectationFailed
	case "500":
		return InternalServerError
	case "501":
		return NotImplemented
	case "502":
		return BadGateway
	case "503":
		return ServiceUnavailable
	case "504":
		return GatewayTimeout
	case "505":
		return HTTPVersionNot
	}

	return BadRequest
}
