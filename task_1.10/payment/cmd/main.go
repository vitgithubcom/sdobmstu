package main

import (
	"fmt"
	"payment/internal/entity"
)

func main() {
	
	processors := []entity.PaymentProcessor{
		entity.CreditCard{
			CardNumber: "1234-5678-9012-3456",
			CardHolder: "Иван Петров",
			CVV:        "123",
		},
		entity.CryptoWallet{
			WalletAddress: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0",
			Currency:      "USDT",
		},
	}
	
	
	fmt.Println("=== Обработка платежей ===\n")
	
	for i, processor := range processors {
		result := processor.Process(100.50)
		fmt.Printf("%d. %s\n", i+1, result)
		fmt.Println()
	}
}