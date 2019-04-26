package prototype

type BMWCar struct {
	name string
}

func (c *BMWCar) GetName() string {
	return c.name
}

func (c *BMWCar) SetName(name string) {
	c.name = name
}

func (c *BMWCar) Clone() *BMWCar {
	return &BMWCar{name: c.name}
}
