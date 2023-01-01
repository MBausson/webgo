package Http

import "strings"

type HttpHeader struct {
	Key   string
	Value string
}

func HeaderFromString(content string) HttpHeader {
	parts := strings.Split(content, ":")
	return HttpHeader{parts[0], strings.Join(parts[1:], ":")[1:]}
}

func (h HttpHeader) ToString() string {
	return h.Key + ": " + h.Value
}
