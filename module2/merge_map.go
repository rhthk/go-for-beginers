package module2

func MergeMap(m1, m2 map[string]int) map[string]int {
	m := make(map[string]int)
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		if x, ok := m[k]; ok {
			m[k] = x + v
		} else {
			m[k] = v
		}
	}

	return m
}
