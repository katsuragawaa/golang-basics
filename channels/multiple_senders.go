package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	multipleSenderReceiver()
	bufferedChannel()
}

func multipleSenderReceiver() {
	println("multiple sender and receiver")
	ch := make(chan int)

	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		j := j // intermediate variable
		go func() {
			ch <- 2 * j
			wg.Done()
		}()
	}

	wg.Wait()
}

func bufferedChannel() {
	println("buffered")
	ch := make(chan int, 50)

	wg.Add(2)
	go func(ch <-chan int) {
		//for i := range ch {
		//	fmt.Println(i)
		//}

		// this syntax it's the same as the for range above
		for {
			if i, ok := <-ch; ok { // if channel is open, ok is true
				fmt.Println(i, ok)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 32
		ch <- 19
		close(ch) // close the channel before deadlock
		//ch <- 20  :: this would panic since channel is closed
		wg.Done()
	}(ch)

	wg.Wait()
}
