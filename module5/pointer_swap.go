package module5

func swapByPointer(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
