package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 2) // buffered go channel
	// ch := make(chan int) // unbuffered go channel
	var wg sync.WaitGroup

	// G1 goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Sending 1 from G1 go routine")
		ch <- 1 // channel theke data send kora hoise, arek channel theke jotokkhon na data ta receive kora hobe totokkhon o sleep mode e thakbe.
	}()

	// G2 goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Sending 2 from G2 go routine")
		ch <- 2
	}()

	// G3 goroutine
	wg.Add(1)
	go func() {
		wg.Done()
		fmt.Println("Receiving data from G3 go routine")
		data := <-ch
		fmt.Println("Data =", data)
	}()
	wg.Wait()
	fmt.Println("Main goRoutine ends")
}

/*
Go Channel Two types: Unbuffered, Buffered

*/