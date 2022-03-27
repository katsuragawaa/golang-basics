package main

import "fmt"

func main() {
	sayMessage("heelu~", 10)
	sameType("a", "b", "c")

	text := "hey"
	passingCopy(text)
	passingPointer(&text)
	fmt.Println(text)

	s := sumPointerReturn(1, 2, 3, 4)
	fmt.Println("sum =", *s)

	r := declareReturn(1, 2)
	fmt.Println(r)

	e, err := errorReturn(2.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e)
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

func sumPointerReturn(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	return &result
}

func declareReturn(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

func errorReturn(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Can't divide by zero")
	}
	return a / b, nil
}
