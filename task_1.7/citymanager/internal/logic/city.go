package logic

import (
	"strings"
)

func CreateInitialCities() []string {
	return []string{}
}

func AddCity(cities []string, city string) []string {
	for _, existing := range cities {
		if strings.EqualFold(existing, city) {
			return cities
		}
	}
	return append(cities, city)
}

func RemoveCity(cities []string, city string) []string {
	index := -1
	for i, existing := range cities {
		if strings.EqualFold(existing, city) {
			index = i
			break
		}
	}
	
	if index == -1 {
		return cities
	}
	
	return append(cities[:index], cities[index+1:]...)
}

func FindCity(cities []string, city string) (found bool, index int) {
	for i, existing := range cities {
		if strings.EqualFold(existing, city) {
			return true, i
		}
	}
	return false, -1
}