package main

import "fmt"

func main() {
	g := greeter{
		greeting: "hellu",
		name:     "john",
	}
	g.greet()
	fmt.Println("greeting:", g.greeting)

	println("---")
	g.overwriteGreet()
	fmt.Println("greeting:", g.greeting)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // operating with a copy
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) overwriteGreet() { // reference original
	fmt.Println(g.greeting, g.name)
	g.greeting = "yo"
}
