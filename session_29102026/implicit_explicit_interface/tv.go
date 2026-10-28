package main

import "fmt"

type Device interface{
	TurnOn()
}

type TV struct{}

func (t TV) TurnOn(){
	fmt.Println("TV is now ON")
}


func main(){
	var d Device = TV{}
	d.TurnOn()

	
}