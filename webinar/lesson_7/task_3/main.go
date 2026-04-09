package main

import "fmt"

func change0index(arr [3]string){
	arr[0] = "different planet"
}


func main() {
	planets := [...] string{
		"Меркурий",
		"Венера",
		"Земля",
	}

change0index(planets)

	fmt.Println(planets[0])

}

