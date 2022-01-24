package ablib

type Component struct {
	Name string
	Dist int
}

// Returns the component's name
func (c *Component) String() string {
	return c.Name
}
