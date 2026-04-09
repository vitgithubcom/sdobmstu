package main

import "fmt"

func changeZeroIndex(arr []int) {
	arr[0] = 99999
}


func main() {
	planets := [] string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
	}
	
	//planets_2 := []string{}
	planets_2 := make([]string,1)
	planets_2[0] = planets[0]

	planets_3 := []string{}
	planets_3 = append (planets_3, planets[0])

	fmt.Println(planets_3)

	planets_4 := []int{}
	for i := 0; i < 30; i++ {
		planets_4 = append(planets_4, i)
	}

	fmt.Println(planets_4)
	fmt.Println(len(planets_4))
	fmt.Println(cap(planets_4))

	city := []string{"Москва","Сочи","Лондон","Череповец"}
	fmt.Println(city)
	city = append(city[:2],city[3:]...)
	fmt.Println(city)
	fmt.Println(len(city))

	planets_5 := make([]int, 0, 10)
	for i := 0; i < 30; i++ {
		planets_5 = append(planets_5, i)
		fmt.Printf("len slice: %v, cap slice: %v\n", len(planets_5), cap(planets_5))
	}

	planets_6 := make([]int, 0)
	for i := 0; i < 30; i++ {
		planets_6 = append(planets_6, i)
		fmt.Printf("len slice: %v, cap slice: %v\n", len(planets_6), cap(planets_6))
	}

	planets_7 := make([]int, 0, 100)
	for i := 0; i < 30; i++ {
		planets_7 = append(planets_7, i)
	}
	changeZeroIndex(planets_7)
	planets_7 = append(planets_7[:1], planets_7[len(planets_7)-1:]...)
	fmt.Println(planets_7)
}

