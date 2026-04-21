package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(index int) {
            fmt.Printf("Горутина %d начала работу\n", index)
            time.Sleep(10 * time.Millisecond) 
            fmt.Printf("Горутина %d закончила\n", index)
            wg.Done()
        }(i)
        fmt.Printf("Создана горутина %d\n", i)
    }
    
    wg.Wait()
}