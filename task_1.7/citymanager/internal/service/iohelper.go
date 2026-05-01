package service

import (
	"fmt"
)

func PrintMenu() {
	fmt.Println("=== Управление списком городов ===")
	fmt.Println("Доступные операции: добавление, удаление, поиск")
	fmt.Println()
}

func PrintCities(cities []string, message string) {
	fmt.Printf("%s:\n", message)
	if len(cities) == 0 {
		fmt.Println("  Список городов пуст")
	} else {
		for i, city := range cities {
			fmt.Printf("  %d. %s\n", i+1, city)
		}
	}
	fmt.Println()
}

func PrintSearchResult(city string, found bool, index int) {
	if found {
		fmt.Printf("Город \"%s\" найден на позиции %d\n", city, index+1)
	} else {
		fmt.Printf("Город \"%s\" не найден в списке\n", city)
	}
	fmt.Println()
}