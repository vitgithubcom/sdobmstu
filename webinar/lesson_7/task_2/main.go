package main

import "fmt"

func main() {
	planets := [...] string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран супер",
	}
	fmt.Println(len(planets))

	for i := 0; i < len(planets); i++ {
		fmt.Println(planets[i])
	}

	for i, v := range planets {
		fmt.Println(i,v)
	}

	for i, v := range planets {		
		fmt.Println(v)
		planets[i] = "bad planet"
	}
	fmt.Println(planets)
}
