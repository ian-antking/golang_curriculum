package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSharePizza (test *testing.T) {
	test.Run("given a number of pizzas, and an even number of people, returns number of slices for each person", func(test *testing.T) {
		const expected = 4
		var actual int = sharePizza(1, 2).slicesPerPerson

		assert.Equal(test, expected, actual)
	})
}