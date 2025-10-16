package module4

type Rectangle struct {
	length  float64
	breadth float64
}

func (rect Rectangle) area() float64 {
	return rect.breadth * rect.length
}

func (rect Rectangle) perimeter() float64 {
	return 2 * (rect.breadth + rect.length)
}
