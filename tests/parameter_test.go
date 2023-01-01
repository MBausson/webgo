package tests

import (
	"testing"
	"webgo/pkg/Http"
)

func TestSimpleParameter(t *testing.T) {
	expected := Http.HttpParameter{
		Name:  "Name",
		Value: "Value",
	}

	got := Http.ParameterFromString("Name=Value")

	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestMultipleEquals(t *testing.T) {
	expected := Http.HttpParameter{
		Name:  "Name",
		Value: "Value=Value1=Value2",
	}

	got := Http.ParameterFromString("Name=Value=Value1=Value2")

	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestSimpleUrl(t *testing.T) {
	got := Http.UrlToParameters("?Name1=Value1&Name2=Value2&Name3=Value3")

	n1, _ := got.Get("Name1")
	n2, _ := got.Get("Name2")
	n3, _ := got.Get("Name3")

	if n1.Value != "Value1" || n2.Value != "Value2" || n3.Value != "Value3" {
		t.Fail()
	}
}
