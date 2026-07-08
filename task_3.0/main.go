package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)


type ApiResponse struct {
	Data []CryptoData `json:"data"`
	Status struct {
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"status"`
}

type CryptoData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  []Quote `json:"quote"`  
}

type Quote struct {
	Id     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

func main() {

	apiKey := os.Getenv("CMC_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ Ошибка: CMC_API_KEY не установлен")
		fmt.Println("💡 Используйте: CMC_API_KEY=your_key go run main.go")
		os.Exit(1)
	}


	cryptoIDs := os.Getenv("CRYPTO_IDS")
	if cryptoIDs == "" {
		cryptoIDs = "1,1027" 
	}


	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v3/cryptocurrency/quotes/latest?id=%s&convert=USD", cryptoIDs)

	fmt.Printf("🔍 Запрос к API для ID: %s\n", cryptoIDs)
	fmt.Println("⏳ Ожидание ответа...")


	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("❌ Ошибка создания запроса: %v\n", err)
		os.Exit(1)
	}


	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.Header.Add("Accept", "application/json")


	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Ошибка выполнения запроса: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Ошибка чтения ответа: %v\n", err)
		os.Exit(1)
	}


	if resp.StatusCode != http.StatusOK {
		fmt.Printf("❌ API вернул ошибку (статус %d): %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}


	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("❌ Ошибка парсинга JSON: %v\n", err)
		fmt.Printf("📄 Полученный ответ: %s\n", string(body))
		os.Exit(1)
	}


	if len(apiResponse.Data) == 0 {
		fmt.Println("⚠️ Данные не найдены")
		os.Exit(1)
	}

	
	fmt.Println("\n📊 Результат:")
	fmt.Println("[")
	for i, crypto := range apiResponse.Data {
		if i > 0 {
			fmt.Println(",")
		}
		
		
		price := 0.0
		if len(crypto.Quote) > 0 {
			price = crypto.Quote[0].Price
		}
		
		fmt.Printf("{\n    name: %s,\n    priceUsd: %.16f\n}", crypto.Name, price)
	}
	fmt.Println("\n]")
	fmt.Printf("\n✅ Успешно получено %d записей\n", len(apiResponse.Data))
}