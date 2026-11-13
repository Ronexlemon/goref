package channel

import (
    "fmt"
    "sync"
)

func producer(s chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()

    for i := 0; i < 10; i++ {
        s <- i
    }

    close(s)
}

func consumer(r <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    sum := 0

    for value := range r {
        sum += value
    }

    fmt.Println("The sum:", sum)
}

func CalculateSum() {
    c := make(chan int)

    var wg sync.WaitGroup

    wg.Add(2)

    go producer(c, &wg)
    go consumer(c, &wg)

    wg.Wait()
}