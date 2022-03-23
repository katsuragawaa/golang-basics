package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i) // convert int to float
	fmt.Printf("%v, %T\n", j, j)

	// you can convert float to int, but it lose the decimal
	var a float32 = 9.9
	fmt.Printf("%v, %T\n", a, a)

	var b int
	b = int(a)
	fmt.Printf("%v, %T\n", b, b)

	// converting integer to string
	var s int = 20
	fmt.Printf("%v, %T\n", s, s)

	var r string
	// wrong way to converto to string
	r = string(s)
	fmt.Printf("%v, %T\n", r, r)

	// correct way using strconv package
	r = strconv.Itoa(s)
	fmt.Printf("%v, %T\n", r, r)
}
