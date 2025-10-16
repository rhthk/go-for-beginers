package module5

func passByValue(n int) int {
	n = n + 2
	return n
}

func passByPointer(n *int) int {
	*n = *n + 3
	return *n
}
