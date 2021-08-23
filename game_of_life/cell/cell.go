package cell

type Cell struct {
	Alive bool
}

func NewCell() Cell {
	return Cell{ Alive: true }
}

func (c *Cell) Die() {
	c.Alive = false
}

func (c *Cell) Live() {
	c.Alive = true
}
