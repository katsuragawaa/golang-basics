package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	a := 1 == 1
	b := 1 == 2
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	// cool fact, if you don't initialize a variable
	// it will initialize as a zero value of that type
	var x bool
	fmt.Printf("%v, %T\n", x, x)
}
