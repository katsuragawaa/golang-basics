package main

import (
	"fmt"
)

func main() {
	var i interface{} = 0

	switch i.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("another type")
	}
}
