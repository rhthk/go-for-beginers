package module4

type Complex struct {
	real float64
	img  float64
}

func (c *Complex) addComplex(n Complex) {
	c.real = c.real + n.real
	c.img = c.img + n.img
}
