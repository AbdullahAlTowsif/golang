package main

import (
	"fmt"
	"time"
)

var a = 10
const p = 11

func printHello(num int) {
	fmt.Println("Hello towsif", num)
}

func main() {
	var x int = 10
	fmt.Println("Hellllllo", x)
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)

	fmt.Println(a, " ", p)

	time.Sleep(5*time.Second) // time.Second = 1s
}
