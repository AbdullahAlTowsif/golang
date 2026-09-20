package main

import (
	"fmt"
	"log"
)

func panicSomething() {
	defer func() {
		fmt.Println("From deferred function")
		r := recover()
		if r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	fmt.Println("Hello from panic something!")

	panic("Something really bad happened!")
}

func fatalSomething() {
	defer func () {
		fmt.Println("From fatal deferred funtion")
	}()

	log.Fatal()
	fmt.Println("Everything is finished!")
}

func main() {
	// panicSomething()
	fatalSomething()
	fmt.Println("From main")
}
