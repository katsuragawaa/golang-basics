package main

import "fmt"

func main() {
	simpleLoop()
	nestedLoops()
	forRangeLoop()
}

func simpleLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	println("\n---")
	for i, j := 0, 0; i < 5 || j < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	println("\n---")
	index := 0
	for index < 2 { // while loop like
		fmt.Println(index)
		index++
	}

	println("\n---")
	for { // infinite loop
		fmt.Println(index)
		if index == 5 {
			break // exit loop
		}
		index++
	}

	println("\n---")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip loop
		}
		fmt.Println(i)
	}
}

func nestedLoops() {
	println("\n---")
Loop: // label, break out to here
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
}

func forRangeLoop() {
	println("\n---")
	slice := []int{1, 2, 3}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	println("\n---")
	hash := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range hash {
		fmt.Println(key, value)
	}

	println("\n---")
	str := "Hellou"
	for index, unicode := range str {
		fmt.Println(index, string(unicode))
	}
}
