package tests

import (
	"testing"
	"webgo/pkg/Http"
)

func TestSimple(t *testing.T) {
	expected := Http.HttpRequest{
		Method:        Http.GET,
		EndpointRoute: "/endpoint",
		CompleteRoute: "/endpoint",
		Version:       "HTTP/1.1",
		Headers:       []Http.HttpHeader{},
		Content:       "",
	}

	got, _ := Http.RequestFromString("GET /endpoint HTTP/1.1\r\n")

	if !expected.Equals(got) {
		t.Fatalf(`Expected %v, got %v`, expected, got)
	}
}

func TestMultiHeaders(t *testing.T) {
	expected := Http.HttpRequest{
		Method:        Http.GET,
		EndpointRoute: "/endpoint",
		CompleteRoute: "/endpoint",
		Version:       "HTTP/1.1",
		Headers:       []Http.HttpHeader{Http.HttpHeader{"Header1", "Value1"}, Http.HttpHeader{"Header2", "Value2"}, Http.HttpHeader{"Header3", "Value3"}},
		Content:       "",
	}

	got, _ := Http.RequestFromString("GET /endpoint HTTP/1.1\r\nHeader1: Value1\r\nHeader2: Value2\r\nHeader3: Value3\r\n")

	if !expected.Equals(got) {
		t.Fatalf(`Expected %v, got %v`, expected, got)
	}
}

func TestContent(t *testing.T) {
	expected := Http.HttpRequest{
		Method:        Http.GET,
		EndpointRoute: "/endpoint",
		CompleteRoute: "/endpoint",
		Version:       "HTTP/1.1",
		Headers:       []Http.HttpHeader{},
		Content:       "Gordon Freeman in the flesh... or rather in the hazardous suit\r\nI took the liberty of relieving of your weapons",
	}

	got, _ := Http.RequestFromString("GET /endpoint HTTP/1.1\r\n\r\nGordon Freeman in the flesh... or rather in the hazardous suit\r\nI took the liberty of relieving of your weapons")

	if !expected.Equals(got) {
		t.Fatalf(`Expected %v, got %v`, expected, got)
	}
}

func TestComplete(t *testing.T) {
	expected := Http.HttpRequest{
		Method:        Http.GET,
		EndpointRoute: "/endpoint",
		CompleteRoute: "/endpoint",
		Version:       "HTTP/1.1",
		Headers:       []Http.HttpHeader{Http.HttpHeader{"Header1", "Value1"}, Http.HttpHeader{"Header2", "Value2"}, Http.HttpHeader{"Header3", "Value3"}},
		Content:       "A request message\r\nWith 2 lines\r\n\r\nWith 3 lines !",
	}

	got, _ := Http.RequestFromString("GET /endpoint HTTP/1.1\r\nHeader1: Value1\r\nHeader2: Value2\r\nHeader3: Value3\r\n\r\nA request message\r\nWith 2 lines\r\n\r\nWith 3 lines !")

	if !expected.Equals(got) {
		t.Fatalf(`Expected %v, got %v`, expected, got)
	}
}

func TestUrlParameters(t *testing.T) {
	expected := Http.HttpRequest{
		Method:        Http.GET,
		EndpointRoute: "/endpoint",
		CompleteRoute: "/endpoint?param=value",
		Version:       "HTTP/1.1",
		Headers:       []Http.HttpHeader{Http.HttpHeader{"Header1", "Value1"}, Http.HttpHeader{"Header2", "Value2"}, Http.HttpHeader{"Header3", "Value3"}},
		Content:       "A request message\r\nWith 2 lines\r\n\r\nWith 3 lines !",
	}

	got, _ := Http.RequestFromString("GET /endpoint?param=value HTTP/1.1\r\nHeader1: Value1\r\nHeader2: Value2\r\nHeader3: Value3\r\n\r\nA request message\r\nWith 2 lines\r\n\r\nWith 3 lines !")

	if !expected.Equals(got) {
		t.Fatalf(`Expected %v, got %v`, expected, got)
	}
}

func TestUnreadableSilly(t *testing.T) {
	_, err := Http.RequestFromString("black mesa")

	if err == nil {
		t.Fail()
	}
}

func TestUnreadableWords(t *testing.T) {
	_, err := Http.RequestFromString("GET /endpoint HTTP 1.1\r\nHeader1: Value1\r\n\r\nMessage body")

	if err == nil {
		t.Fail()
	}
}
