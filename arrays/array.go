package main

import "fmt"

func main() {
	// variable := [length]type{initialize the array}
	grades := [3]int{60, 70, 70}
	fmt.Printf("Grades: %v\n", grades)

	a := [...]int{2, 3, 4}
	fmt.Printf("Grades: %v\n", a)
	fmt.Printf("Array length: %v\n\n", len(a))

	var b [3]int
	fmt.Printf("Empty array with length of 3: %v\n\n", b)

	matrix := [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Printf("%v, %T\n\n", matrix, matrix)

	original := [...]int{1, 2, 3}
	new := original // create a new array, isn't a pointer
	// new := &original // using pointer
	new[1] = 999
	fmt.Println(original)
	fmt.Println(new)
}
