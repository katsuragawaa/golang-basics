package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go sayHello()
	badPractice()
	goodPractice()
	wg.Wait() // go will wait all the routines added to the group before exit function
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}

func badPractice() {
	msg := "bad hey"
	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	// race condition, reassign variable before routine print, so it'll say "bye"
	// go run -race to check
	msg = "bad bye"
}

func goodPractice() {
	msg := "good hey"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "good bye"
}
