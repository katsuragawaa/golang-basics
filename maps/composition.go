package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal // embedding
	Speed  float32
	CanFly bool
}

func main() {
	emu := Bird{}
	emu.Name = "Emu"
	emu.Origin = "AUS"
	emu.Speed = 34
	emu.CanFly = false
	fmt.Println(emu)
	fmt.Println(emu.Name)

	crow := Bird{
		Animal: Animal{Name: "Crow", Origin: "Unknown"},
		Speed:  60,
		CanFly: true,
	}
	fmt.Println(crow)
	fmt.Println(crow.Origin)
}
