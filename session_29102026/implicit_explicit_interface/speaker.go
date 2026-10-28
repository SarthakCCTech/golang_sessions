package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}

type Robot struct{}

func (r Robot) Speak() {
    fmt.Println("Hello!")
}

func main() {
    var s Speaker

    s = Dog{}
    s.Speak()

    s = Robot{}
    s.Speak()
}		