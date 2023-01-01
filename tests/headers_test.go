package tests

import (
	"testing"
	"webgo/pkg/Http"
)

func TestValid(t *testing.T) {
	expected := Http.HttpHeader{"Keey", "Vaalue"}
	got := Http.HeaderFromString("Keey: Vaalue")

	if expected != got {
		t.Fatalf(`Expected %v, got %v`, expected, got)
	}
}

func TestMultiColon(t *testing.T) {
	expected := Http.HttpHeader{"Key", "Value1:Value2:Value3"}
	got := Http.HeaderFromString("Key: Value1:Value2:Value3")

	if expected != got {
		t.Fatalf(`Expected %v, got %v`, expected, got)
	}
	
}
