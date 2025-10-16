package module4

type Student struct {
	name  string
	mark  int
	grade string
}

func (s *Student) updateMark(mark int) {
	s.mark = mark
	switch {
	case mark > 90:
		s.grade = "A"
	case mark > 80:
		s.grade = "B"
	case mark > 70:
		s.grade = "C"
	case mark > 60:
		s.grade = "D"
	default:
		s.grade = "F"
	}
}
