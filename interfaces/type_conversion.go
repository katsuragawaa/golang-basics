package main

import (
	"fmt"
)

func main() {
	var i interface{} = 0
	// in go 1.8 "any" is an alias for empty interface

	switch i.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("another type")
	}
}
