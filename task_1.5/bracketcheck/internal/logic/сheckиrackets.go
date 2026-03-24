package logic

// CheckBrackets проверяет правильность расстановки круглых скобок
func CheckBrackets(s string) (openCount, closeCount int, isValid bool) {
	openCount = 0
	closeCount = 0
	counter := 0
	
	for _, char := range s {
		if char == '(' {
			openCount++
			counter++
		} else if char == ')' {
			closeCount++
			counter--
		}
		
		// Если закрывающих скобок стало больше, чем открывающих
		if counter < 0 {
			return openCount, closeCount, false
		}
	}
	
	// Скобки расставлены правильно, если счетчик вернулся к нулю
	isValid = (counter == 0)
	return openCount, closeCount, isValid
}