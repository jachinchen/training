package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("apple")
	if hello != "Hi, apple" {
		t.Errorf("Testing fail")
	}

	hello = HelloWorld("apple ")
	if hello != "Hi, apple " {
		t.Errorf("Testing fail")
	}
}
