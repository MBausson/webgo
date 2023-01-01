package Http

type UnreadableRequest struct{}

func (err *UnreadableRequest) Error() string {
	return "Couldn't read HTTP request. Is it really at HTTP standards ?"
}
