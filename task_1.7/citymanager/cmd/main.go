package main

import (
	"citymanager/internal/logic"
	"citymanager/internal/service"
)

func main() {
	
	cities := logic.CreateInitialCities()
	
	
	service.PrintMenu()
	
	
	cities = logic.AddCity(cities, "Москва")
	cities = logic.AddCity(cities, "Санкт-Петербург")
	cities = logic.AddCity(cities, "Новосибирск")
	
	service.PrintCities(cities, "После добавления городов")
	
	
	cities = logic.RemoveCity(cities, "Москва")
	service.PrintCities(cities, "После удаления города Москва")
	
	
	searchCity := "Новосибирск"
	found, index := logic.FindCity(cities, searchCity)
	service.PrintSearchResult(searchCity, found, index)
	
	
	searchCity = "Владивосток"
	found, index = logic.FindCity(cities, searchCity)
	service.PrintSearchResult(searchCity, found, index)
}