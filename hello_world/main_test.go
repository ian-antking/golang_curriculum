package main

import "testing"

func TestHello(t *testing.T) {
	const expected = "Hello World!"
	var actual string = hello();

	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}