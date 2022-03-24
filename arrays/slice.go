package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[1] = 9 // it's different from array, here it's a pointer
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n\n", cap(a))

	x := []int{1, 2, 3, 4, 5, 6}
	y := x[:]
	c := x[3:]
	v := x[:5]
	n := x[1:4]

	n[2] = 999 // still with reference, mutate original

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(c)
	fmt.Println(v)
	fmt.Println(n)

	println("----")
	// make
	makeFunc()

	println("----")
	// append
	appending()

	println("----")
	// removing elements
	removing()
}

func makeFunc() {
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n\n", cap(a))
}

func appending() {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n\n", cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n\n", cap(a))

	a = append(a, 1, 2, 3, 4)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))

	b := []int{5, 6}
	a = append(a, b...) // append 5 and 6 to the end of the slice
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
}

func removing() {
	a := []int{1, 2, 3, 4}
	b := a[1:]
	c := a[:len(a)-1]
	fmt.Println("original slice: ", a)
	fmt.Println("without first elem: ", b)
	fmt.Println("without last elem: ", c)

	d := append(a[:2], a[3:]...)
	fmt.Println("without index 2: ", d)
}
