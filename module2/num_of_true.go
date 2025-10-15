package module2

func CountOfTrue(slice []bool) int {
	count := 0
	for _, v := range slice {
		if v {
			count++
		}
	}
	return count
}
