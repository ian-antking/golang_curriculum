package cell

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCell(t *testing.T) {
	t.Run("cells have a life property which is set to true by default", func(t *testing.T) {
		cell := NewCell()

		assert.True(t, cell.Alive)
	})

	t.Run("cells can die", func(t *testing.T) {
		cell := NewCell()

		cell.Die()

		assert.False(t, cell.Alive)
	})

	t.Run("cells can come back to life", func(t *testing.T) {
		cell := NewCell()

		cell.Live()

		assert.True(t, cell.Alive)
	})

	t.Run("cell methods only affect individual instances", func(t *testing.T) {
		cell1 := NewCell()
		cell2 := NewCell()

		cell1.Die()

		assert.False(t, cell1.Alive)
		assert.True(t, cell2.Alive)
	})
}