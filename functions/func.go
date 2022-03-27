package main

import "fmt"

func main() {
	sayMessage("heelu~", 10)
	sameType("a", "b", "c")

	text := "hey"
	passingCopy(text)
	passingPointer(&text)
	fmt.Println(text)
}

func sayMessage(message string, index int) {
	fmt.Println(message, index)
}

func sameType(a, b, c string) { // syntax sugar
	fmt.Printf("\n%T, %T, %T\n", a, b, c)
}

func passingCopy(text string) {
	text = "safe to change"
	fmt.Println(text)
}

func passingPointer(text *string) {
	*text = "change the original"
	fmt.Println(*text)
}
