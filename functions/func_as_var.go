package main

import "fmt"

func main() {
	// anonymous function
	func() {
		msg := "hi~"
		fmt.Println(msg)
	}()

	for i := 0; i < 2; i++ {
		// this works but it's dangerous if function is async
		func() {
			fmt.Println(i)
		}()
		// best practice
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	f := func() { // or var f func() = func() {...}
		fmt.Println("function as variable")
	}
	f()

	passingFunctions()
}

func passingFunctions() {
	var divide func(float64, float64) (float64, error)
	divide = func(f1, f2 float64) (float64, error) {
		if f2 == 0.0 {
			return 0.0, fmt.Errorf("can't divide by zero")
		} else {
			return f1 / f2, nil
		}
	}

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
