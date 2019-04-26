package prototype

type BenChiCar struct {
	name string
}

func (c *BenChiCar) GetName() string {
	return c.name
}

func (c *BenChiCar) SetName(name string) {
	c.name = name
}

func (c *BenChiCar) Clone() *BenChiCar {
	return &BenChiCar{name: c.name}
}
