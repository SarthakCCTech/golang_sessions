package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Dog struct{}

func (d Dog) Eat(){
	fmt.Println("Dog is eating")
}

func (d Dog) Sleep(){
	fmt.Println("Dog is sleeping")
}

func main(){
	var a Animal = Dog{}

	a.Eat()
	a.Sleep()
}