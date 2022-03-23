package main

import "fmt"

const a = iota // 0, int

const (
	b = iota // 0
	c = iota // 1
	d = iota // 2
)

// auto inherit, mostly used for enums
const (
	x = iota
	y
	z
)

func main() {
	println("a")
	fmt.Printf("%v\n", a)
	println("b to d")
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	println("x to z")
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}
