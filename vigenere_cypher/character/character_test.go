package character

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Run("given a character a-z and an integer of 1, returns shifted character", func(t *testing.T) {
		const expected = 'b'
		actual, _ := Shift('a', 1)

		assert.Equal(t, expected, actual)
	})

	t.Run("can shift a character multiple times", func(t *testing.T) {
		const expected = 'e'
		actual, _ := Shift('a', 4)

		assert.Equal(t, expected, actual)
	})

	t.Run("will reset to 'a' if shifted past 'z'", func(t *testing.T) {
		const expected = 'a'
		actual, _ := Shift('z', 1)

		assert.Equal(t, expected, actual)
	})

	t.Run("will convert uppercase letters to lowercase", func(t *testing.T) {
		const expected = 'a'
		actual, _ := Shift('A', 0)

		assert.Equal(t, expected, actual)
	})

	t.Run("will throw if given a non alpha character", func(t *testing.T) {
		_, err := Shift('!', 3)

		assert.Error(t, err)
	})
}