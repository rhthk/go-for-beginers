package module2

func SwapVar(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
