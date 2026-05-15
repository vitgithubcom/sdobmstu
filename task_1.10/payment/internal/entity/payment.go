package entity

import (
	"fmt"
)


type PaymentProcessor interface {
	Process(amount float64) string
}


type CreditCard struct {
	CardNumber string
	CardHolder string
	CVV        string
}


func (c CreditCard) Process(amount float64) string {
	
	maskedNumber := c.CardNumber[len(c.CardNumber)-4:]
	return fmt.Sprintf("ok Оплата %.2f руб. с карты ****%s (владелец: %s) прошла успешно", 
		amount, maskedNumber, c.CardHolder)
}


type CryptoWallet struct {
	WalletAddress string
	Currency      string
}

func (cw CryptoWallet) Process(amount float64) string {
	shortAddress := cw.WalletAddress[:10] + "..."
	return fmt.Sprintf("ok Перевод %.2f %s на кошелек %s выполнен успешно", 
		amount, cw.Currency, shortAddress)
}