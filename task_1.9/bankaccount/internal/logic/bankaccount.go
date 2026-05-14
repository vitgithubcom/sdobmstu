package logic

import (
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма снятия должна быть положительной")
		return
	}
	
	if amount <= b.Balance {
		b.Balance -= amount
	} else {
		fmt.Printf("Ошибка: недостаточно средств! Доступно: %.2f руб., запрошено: %.2f руб.\n", b.Balance, amount)
	}
}