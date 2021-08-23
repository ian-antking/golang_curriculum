package cell

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCell(t *testing.T) {
	t.Run("cells have a life property which is set to true by default", func(t *testing.T) {
		cell := NewCell(0, 0)

		assert.True(t, cell.Alive)
	})

	t.Run("cells can die", func(t *testing.T) {
		cell := NewCell(0, 0)

		cell.Die()

		assert.False(t, cell.Alive)
	})

	t.Run("cells can come back to life", func(t *testing.T) {
		cell := NewCell(0, 0)

		cell.Live()

		assert.True(t, cell.Alive)
	})

	t.Run("cells have x and y coordinates that are set at instantiation", func(t *testing.T) {
		cell := NewCell(1, 1)

		assert.Equal(t, 1, cell.X)
		assert.Equal(t, 1, cell.Y)
	})

	t.Run("cell methods only affect individual instances", func(t *testing.T) {
		cell1 := NewCell(0, 0)
		cell2 := NewCell(0, 0)

		cell1.Die()

		assert.False(t, cell1.Alive)
		assert.True(t, cell2.Alive)
	})
}