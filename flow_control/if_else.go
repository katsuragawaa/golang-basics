package main

import "fmt"

func main() {
	basic()
	println("---")
	initializerSyntax()
	println("---")
	comparison()
}

func basic() {
	if true {
		fmt.Println("Test is true")
	}
}

func initializerSyntax() {
	statePopulations := map[string]int{
		"California": 40000000,
		"New York":   20000000,
		"Ohio":       10000000,
	}

	if population, ok := statePopulations["Ohio"]; ok {
		fmt.Println(population)
	}
}

func comparison() {
	number := 50
	guess := 50

	if guess < 1 {
		fmt.Println("Guess need to be greater than 1")
	} else if guess > 100 {
		fmt.Println("Less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("Nice")
		}
	}
}
