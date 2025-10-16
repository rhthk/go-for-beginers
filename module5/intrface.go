package module5

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return 3.1415 * c.r * c.r
}

func (c Circle) perimeter() float64 {
	return c.r * 2 * 3.1415
}
