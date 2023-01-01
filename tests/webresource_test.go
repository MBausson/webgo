package tests

import (
	"testing"
	"webgo/pkg/Resources"
)

func TestGetResult(t *testing.T) {
	expected := "A simple text file..."
	wr := Resources.WebResource{
		LocalPath: "./test_file.txt",
	}

	wr.Load()
	got := wr.GetResult()

	if expected != got {
		t.Failed()
	}
}
