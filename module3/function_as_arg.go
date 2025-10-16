package module3

func add(numbs []int) int {
	sum := 0
	for _, v := range numbs {
		sum += v
	}
	return sum
}

func multiply(numbs []int) int {
	product := 1
	for _, v := range numbs {
		product *= v
	}
	return product
}

func calculate(fn func([]int) int, numbs ...int) int {
	return fn(numbs)
}
