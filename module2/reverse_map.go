package module2

func ReverseMap(input map[string]int) map[int]string {
	out := make(map[int]string)
	for k, v := range input {
		out[v] = k
	}
	return out
}
