package main

import "fmt"

func main() {
	s := sum("hi", 1, 2, 3, 4, 5)
	fmt.Println(s)
}

func sum(message string, values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(message)
	return result
}
