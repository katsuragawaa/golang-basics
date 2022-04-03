package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // channel messages are strongly typed

	wg.Add(2)

	// receiver
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	// sender, order don't matter?
	go func() {
		i := 42
		ch <- i
		i = 25
		wg.Done()
	}()

	wg.Wait()

	println("----")
	senderReceiver()

	println("----")
	betterSenderAndReceiver()
}

// routines can act as sender and receiver, but not recommended
func senderReceiver() {
	ch := make(chan int)

	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 24
		wg.Done()
	}()

	go func() {
		i := 42
		ch <- i
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()
}

func betterSenderAndReceiver() {
	ch := make(chan int)

	wg.Add(2)

	// receive only
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// send only
	go func(ch chan<- int) {
		i := 42
		ch <- i
		wg.Done()
	}(ch)

	wg.Wait()
}
