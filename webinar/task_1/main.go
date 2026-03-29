package main

import "fmt"

func CheckAge(age int) (c string, d bool) {
	if age >= 18 {
		c = "Можете проходить в наш бар"	
		d = true
	} else if age > 0 && age < 18 {
		c = "Тебе пока рано"
		d = false
	} else {		
		c = "Вы врете"
		d = false
	}	
	return c, d
}

func main() {	
var name string
var age int
fmt.Print("Введите имя:")
fmt.Scanln(&name)
fmt.Print("Введите возраст:")
fmt.Scanln(&age)
message, allowed := CheckAge(age)
	if allowed {
		fmt.Printf("%s! %s\n", message, name)
	} else {
		fmt.Printf("%s\n", message)
	}
}