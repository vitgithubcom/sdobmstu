package logic

import (
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль невозможно")
	}
	return a / b, nil
}