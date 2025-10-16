package module3

func makeMultiplier(n int) func(int) int {
	return func(i int) int {
		return i * n
	}
}
