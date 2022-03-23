package main

import "fmt"

func main() {
	f := 3.14 // gonna be a float64
	fmt.Printf("%v, %T\n", f, f)

	f = 3.8e32
	fmt.Printf("%v, %T\n", f, f)

	e := 33.3
	f = 11.1

	fmt.Println(e + f)
	fmt.Println(e - f)
	fmt.Println(e * f)
	fmt.Println(e / f) // 3 as float
	// fmt.Println(e % f) <- not available
}
