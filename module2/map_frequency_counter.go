package module2

func MapFreqCounter(str string) map[string]int {
	counter := make(map[string]int)
	for _, ch := range str {
		if _, ok := counter[string(ch)]; ok {
			counter[string(ch)]++
		} else {
			counter[string(ch)] = 1
		}
	}
	return counter
}
