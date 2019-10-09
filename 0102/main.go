package main

import (
	"fmt"
)

// реализация интерфейсов

// Vehicle interface
type Vehicle interface {
	move()
}

// Car type
type Car struct {
	model string
}

// Aircraft type
type Aircraft struct {
	model string
}

func (c Car) move() {
	fmt.Println(c.model, "It is car")
}

func (a Aircraft) move() {
	fmt.Println(a.model, "It is aircraft")
}

func main() {
	vaz := Car{model: "2107"}
	tu := Aircraft{model: "140"}
	vehicles := [...]Vehicle{vaz, tu}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
}
