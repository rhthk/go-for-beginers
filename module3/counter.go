package module3

func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}
