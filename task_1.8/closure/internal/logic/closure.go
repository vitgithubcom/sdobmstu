package logic

func CreateCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}