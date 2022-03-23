package main

import "fmt"

func main() {
	var n complex64 = 1 + 2i // imaginary number
	fmt.Printf("%v, %T", n, n)

	a := 1 + 2i
	b := 2 + 5.3i
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	// extract real or imaginary part
	fmt.Println(real(n))
	fmt.Println(imag(n))

	var i complex128 = complex(5, 12) // 5 + 12i
	fmt.Printf("%v, %T", i, i)
}
