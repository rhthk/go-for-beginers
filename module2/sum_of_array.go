package module2

func SumOfArray(slice []int) int {
	var sum int = 0
	for i := range slice {
		sum += slice[i]
	}
	return sum
}
