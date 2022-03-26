package main

import "fmt"

func main() {
	number := 1
	switch number {
	case 1, 3:
		fmt.Println("one or three")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("nether")
	}

	switch i := 1 + 3; i {
	case 1, 3:
		fmt.Println("one or three")
	case 2, 4:
		fmt.Println("two or four")
	default:
		fmt.Println("other")
	}

	switch {
	case number <= 1:
		fmt.Println("less then 1")
		fallthrough // it'll ignore next validation and just run it
	case number <= 5:
		fmt.Println("less then 5")
	default:
		fmt.Println("bigger then 5")
	}

	// type switching
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("integer")
		break
		fmt.Println("this wont run")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("another type")
	}
}
