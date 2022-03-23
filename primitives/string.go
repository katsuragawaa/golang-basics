package main

import "fmt"

func main() {
	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[8], s[8]) // each char it is a uint8 unicode (byte)

	r := " other string"
	fmt.Println(s + r)

	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	a := 'a' // rune
	// var a rune = 'a'
	fmt.Printf("%v, %T", a, a)
}
