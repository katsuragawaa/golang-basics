package main

import "fmt"

func main() {
	types()
	operations()
	operationWithDiffTypes()
	bitOperators()
	bitShifting()
}

func types() {
	println("Types")
	i := 95
	fmt.Printf("%v, %T\n", i, i)

	var u uint16 = 50
	fmt.Printf("%v, %T\n", u, u)

	var b byte = 255 // same as uint8
	fmt.Printf("%v, %T\n", b, b)
}

func operations() {
	println("\nOperations")
	x := 10
	y := 3
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y) // int divided by int will result in a int
	fmt.Println(x % y)
}

func operationWithDiffTypes() {
	println("\nOperations with different types")
	var e int = 10
	var f int8 = 3
	// fmt.Println(e + f) <- this won't work
	intF := int(f)
	fmt.Println(e + intF) // need to convert to same type
}

func bitOperators() {
	println("\nBit operators")
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010 = 2 (both need to have the bit in the position)
	fmt.Println(a | b)  // 1011 = 11 (one, or both, need to have)
	fmt.Println(a ^ b)  // 1001 = 9 (just one, not both)
	fmt.Println(a &^ b) // 0100 = 8 (both don't have)
}

func bitShifting() {
	println("\nBit shift")
	a := 8              // 2³
	fmt.Println(a << 2) // 2³ * 2² = 2⁵ = 32
	fmt.Println(a >> 3) // 2³ / 2³ = 2⁰ = 1
}
