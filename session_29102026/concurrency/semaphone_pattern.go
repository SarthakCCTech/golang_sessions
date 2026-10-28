
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    parking := make(chan struct{}, 2)

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

        go func(car int) {
            defer wg.Done()

            // Take a parking space
            parking <- struct{}{}

            fmt.Println("Car", car, "parked")

            // Stay parked for 2 seconds
            time.Sleep(2 * time.Second)

            fmt.Println("Car", car, "leaving")

            // Free the parking space
            <-parking
        }(i)
    }

    wg.Wait()
    fmt.Println("All cars have finished.")
}
