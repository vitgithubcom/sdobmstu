package main

import "fmt"

func main() {
    
    const pricePerMinute = 12.5
    const pricePerKm = 8.3

    var minutes, kilometers uint

    fmt.Scan(&minutes)
    fmt.Scan(&kilometers)

    costByTime := float64(minutes) * pricePerMinute
    costByDistance := float64(kilometers) * pricePerKm

    fmt.Printf("\nЦена в минутах: %.2f руб.\n", costByTime)
    fmt.Printf("Цена в километрах: %.2f руб.\n", costByDistance)

    if costByTime < costByDistance {
        fmt.Println("Лучше по времени")
    } else if costByDistance < costByTime {
        fmt.Println("Лучше по километражу")
    } else {
        fmt.Println("Оба одинаковы")
    }
}