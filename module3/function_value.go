package module3

func sum() func(int, int) int {
	s := func(x, y int) int {
		return x + y
	}
	return s
}
