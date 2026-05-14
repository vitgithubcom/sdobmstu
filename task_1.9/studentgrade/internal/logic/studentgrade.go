package logic

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}