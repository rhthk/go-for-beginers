package module2

func MaxInSlice(slice []float64) float64 {
	var max float64 = 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}
