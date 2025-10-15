package module2

func StringReverse(str string) string {
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}
