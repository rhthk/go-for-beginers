package module2

func SliceAppend(slice []int, num ...int) []int {
	totalLen := len(slice) + len(num)
	lenOfSlice := len(slice)
	out := make([]int, totalLen)
	for i, v := range slice {
		out[i] = v
	}
	for i, v := range num {
		out[lenOfSlice+i] = v
	}
	return out
}
