package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    balance := 100

    mu.Lock()
    balance = balance + 50
    mu.Unlock()

    fmt.Println(balance)
}