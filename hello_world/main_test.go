package main

import "testing"

func assertCorrectMessage(test testing.TB, expected string, actual string) {
	test.Helper()
	if expected != actual {
		test.Errorf("expected %q but got %q", expected, actual)
	}
}

func TestHello(test *testing.T) {

	test.Run("given no arguments, returns 'Hello World!'", func(test *testing.T) {
		const expected = "Hello World!"
		var actual string = hello();

		assertCorrectMessage(test, expected, actual)
	})

	test.Run("given a string as an argument, returns 'Hello STRING!'", func(test *testing.T) {
		const expected = "Hello Ian!"
		var actual string = hello("Ian")

		assertCorrectMessage(test, expected, actual)
	})

	test.Run("given multiple strings as arguments, returns 'Hello FIRST_STRING!'", func(test *testing.T) {
		const expected = "Hello Dave!"
		var actual string = hello("Dave", "Ian", "James")

		assertCorrectMessage(test, expected, actual)
	})
}