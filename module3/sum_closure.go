package module3

func sumClosure() func(int) int {
	sum := 0
	return func(n int) int {
		sum += n
		return sum
	}
}
