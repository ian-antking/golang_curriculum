package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSharePizza (test *testing.T) {
	test.Run("given one pizza, and an even number of people, returns the number of slices for each person", func(test *testing.T) {
		var expected pizzaReport = pizzaReport{ slicesPerPerson: 4 }
		var actual pizzaReport = sharePizza(1, 2)

		assert.Equal(test, expected, actual)
	})

	test.Run("given multiple pizzas, and an even number of people, returns the number fo slices for each person", func(test *testing.T) {
		var expected pizzaReport = pizzaReport{ slicesPerPerson: 6 }
		var actual pizzaReport = sharePizza(3, 4)

		assert.Equal(test, expected, actual)
	})

	test.Run("given a number of pizzas, and an odd number of people, retuns slicesPerPerson and leftoverSlices", func(test *testing.T) {
		var expected pizzaReport = pizzaReport{ slicesPerPerson: 2,  leftoverSlices: 2}
		var actual pizzaReport = sharePizza(1, 3)

		assert.Equal(test, expected, actual)
	})
}