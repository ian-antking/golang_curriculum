package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSharePizza (test *testing.T) {
	test.Run("given one pizza, and an even number of people, returns the number of slices for each person", func(test *testing.T) {
		var expected int = 4
		actual, _, _ := sharePizza(1, 2)

		assert.Equal(test, expected, actual)
	})

	test.Run("given multiple pizzas, and an even number of people, returns the number fo slices for each person", func(test *testing.T) {
		var expected int = 6
		actual, _, _ := sharePizza(3, 4)

		assert.Equal(test, expected, actual)
	})

	test.Run("given a number of pizzas, and an odd number of people, retuns slicesPerPerson and leftoverSlices", func(test *testing.T) {
		var expectedSlicesPerPerson int = 2
		var expectedLeftoverSlices int = 2
		actualSlicesPerPersn, actualLeftoverSlices, _ := sharePizza(1, 3)

		assert.Equal(test, expectedSlicesPerPerson, actualSlicesPerPersn)
		assert.Equal(test, expectedLeftoverSlices, actualLeftoverSlices)
	})

	test.Run("given 0 pizzas, throws an error", func(test *testing.T) {
		const expectedErrorMessage = "cannot share 0 pizzas among 1 people!"
		_, _, err := sharePizza(0, 1)
		assert.Error(test, err)
		assert.Equal(test, expectedErrorMessage, fmt.Sprint(err))
	})


	test.Run("given 0 people, throws an error", func(test *testing.T) {
		const expectedErrorMessage = "cannot share 1 pizzas among 0 people!"
		_, _, err := sharePizza(1, 0)
		assert.Error(test, err)
		assert.Equal(test, expectedErrorMessage, fmt.Sprint(err))
	})
}