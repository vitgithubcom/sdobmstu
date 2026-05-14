package main

import (
	"fmt"
	"bankaccount/internal/logic"
)

func main() {
	
	account := logic.BankAccount{
		Owner:   "Анна Иванова",
		Balance: 1000.0,
	}
	
	fmt.Printf("Владелец: %s\n", account.Owner)
	fmt.Printf("Начальный баланс: %.2f руб.\n", account.Balance)
	fmt.Println()
	
	account.Deposit(500.0)
	fmt.Printf("После пополнения: %.2f руб.\n", account.Balance)
	
	account.Withdraw(300.0)
	fmt.Printf("После снятия 300 руб.: %.2f руб.\n", account.Balance)
	
	account.Withdraw(1500.0)
	fmt.Printf("Итоговый баланс: %.2f руб.\n", account.Balance)
}