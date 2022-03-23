package main

import "fmt"

// At the package level you need to declare explicitly the type
var h int = 20

// var bloc
var actorName string = "Elisabeth Sladen"
var companion string = "Sarah Jane Smith"
var doctorNumber int = 3
var season int = 11

// cleaner way
var (
	name    string = "Katsuragawa"
	partner string = "Evelyn"
	number  int    = 10
)

func main() {
	// Different formats to declare
	var i int
	i = 42

	var j int = 24

	k := 12

	fmt.Println(i, j, k, h)

	var l float32 = 24 // It would declare as int if we use :=

	fmt.Printf("%v, %T\n", j, j) // value and type
	fmt.Printf("%v, %T", l, l)
}
