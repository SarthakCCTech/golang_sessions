package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    start := make(chan struct{})

    for i := 1; i <= 3; i++ {
        wg.Add(1)

        go func(id int) {
            defer wg.Done()

            fmt.Println("Friend", id, "arrived")

            <-start // Wait at barrier

            fmt.Println("Friend", id, "started trip")
        }(i)
    }

    // Wait for everyone to arrive before opening
    // the barrier (see note below).
    close(start)

    wg.Wait()
}