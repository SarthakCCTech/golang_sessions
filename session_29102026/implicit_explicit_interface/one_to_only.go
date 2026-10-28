package main

import "fmt"

type Device interface {
    TurnOn()
}

type TV struct{}
type Fan struct{}

func (t TV) TurnOn() {
    fmt.Println("TV is ON")
}

func (f Fan) TurnOn() {
    fmt.Println("Fan is ON")
}

func main() {
    var d Device

    d = TV{}
    d.TurnOn()

    d = Fan{}
    d.TurnOn()
}