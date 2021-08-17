package main

import "testing"

func TestHello(test *testing.T) {
	test.Run("given no arguments, returns 'Hello World!'", func(test *testing.T) {
		const expected = "Hello World!"
		var actual string = hello();

		if expected != actual {
			test.Errorf("expected %q but got %q", expected, actual)
		}
	})

	test.Run("given a string as an argument, returns 'Hello STRING!'", func(test *testing.T) {
		const expected = "Hello Ian!"
		var actual string = hello("Ian")

		if expected != actual {
			test.Errorf("expected %q but got %q", expected, actual)
		}
	})

	test.Run("given multiple strings as arguments, returns 'Hello FIRST_STRING!'", func(test *testing.T) {
		const expected = "Hello Dave!"
		var actual string = hello("Dave", "Ian", "James")

		if expected != actual {
			test.Errorf("expected %q but got %q", expected, actual)
		}	
	})
}