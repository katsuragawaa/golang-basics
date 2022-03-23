package main

import "fmt"

var i int = 27
var j int = 2        // <- declared but not used: without error in package level
var Integer int = 99 // titlecase variable can be access outside the package

func main() {
	fmt.Println(i) // 27

	var i int = 42
	// i := 20 <- can't have same variable in the same scope
	fmt.Println(i) // 42

	// j := 20 <- declared but not used: compile error
}
