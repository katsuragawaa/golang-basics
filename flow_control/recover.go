package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		// so we can use defer as a way to recover
		// only this function will stop
		// the others higher in the call stack will keep running
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// panic(err) <- if we need to panic after recover and stop application
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
