package main

import "fmt"

func main() {
	typedConst()
	notTypedConst()
}

func typedConst() {
	const myConst int = 42 // need to be initialized
	// const wrongConst float64 = math.Sin(2) <- can't be defined at runtime
	// myConst = 24 <- nope, it's a constant
	fmt.Printf("%v, %T\n", myConst, myConst)
}

func notTypedConst() {
	const a = 9
	fmt.Printf("%v, %T\n", a, a)

	var b int16 = 1
	fmt.Printf("%v, %T\n", a+b, a+b)

	// wait, but they are different types
	// actually since "a" it's a const without a type declaration
	// the compiler will replace where we use "a" with it's value
	fmt.Printf("%v, %T\n", 9+b, 9+b) // this is what happening
}
