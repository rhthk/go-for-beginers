package module2

func SumOfEven(n int) int {
	sum := 0
	for i := range n {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}
