package module2

func GradeEvaluator(mark int) string {
	switch {
	case mark > 90:
		return "A"
	case mark > 80:
		return "B"
	case mark > 70:
		return "C"
	case mark > 60:
		return "D"
	default:
		return "F"
	}
}
