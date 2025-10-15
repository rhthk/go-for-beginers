package module2

func Fibonacci(n int) []int {
	series := make([]int, n)
	series[1] = 1
	series[2] = 1
	for i := range n {
		if i != 0 && i != 1 {
			series[i] = series[i-1] + series[i-2]
		}
	}
	return series
}
