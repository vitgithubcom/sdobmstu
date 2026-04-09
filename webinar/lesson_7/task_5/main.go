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
		"Уран",
	}

	FirstPlanets := planets[0:3]
	fmt.Println(FirstPlanets)
	SecondPlanets := planets[3:6]
	fmt.Println(SecondPlanets)
	allPlanets := planets[:]
	fmt.Println(allPlanets)
	FirstPlanets[0] = "different"
	SecondPlanets[0] = "different planet"
	fmt.Println(planets)
	fmt.Println(allPlanets)

}

