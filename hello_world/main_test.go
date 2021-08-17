package main

import "testing"

func TestHello_GivenNoArguments_ReturnsHelloWorld(t *testing.T) {
	const expected = "Hello World!"
	var actual string = hello();

	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func TestHello_GivenName_ReturnsHelloName(t *testing.T) {
	const expected = "Hello Ian!"
	var actual string = hello("Ian")

	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func TestHello_GivenMultipleNames_ReturnsHelloFirstName(t *testing.T) {
	const expected = "Hello Dave!"
	var actual string = hello("Dave", "Ian", "James")

	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}