package cell

type Cell struct {
	Alive bool
	X int
	Y int
}

func NewCell(x, y int) Cell {
	return Cell{ Alive: true, X: x, Y: y }
}

func (c *Cell) Die() {
	c.Alive = false
}

func (c *Cell) Live() {
	c.Alive = true
}
