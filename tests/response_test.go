package tests

import (
	"testing"
	"webgo/pkg/Http"
)

func TestSimpleResponse(t *testing.T) {
	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/html"
	got := Http.HttpResponse{
		Version: "HTTP/1.1",
		Code:    Http.OK,
		Headers: []Http.HttpHeader{{"Content-Type", "text/html"}},
	}.ToString()

	if expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestCompleteResponse(t *testing.T) {
	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nHeader2: Value2\r\n\r\n<html><head><title>Barney</title></head><body><p>92i</p></body></html>"
	got := Http.HttpResponse{
		Version: "HTTP/1.1",
		Code:    Http.OK,
		Headers: []Http.HttpHeader{{"Content-Type", "text/html"}, {"Header2", "Value2"}},
		Content: "<html><head><title>Barney</title></head><body><p>92i</p></body></html>",
	}.ToString()

	if expected != got {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}
