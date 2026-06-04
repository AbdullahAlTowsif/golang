package main

import "fmt"

var a = 10

// standard function / named function
func add(x int, y int) {
	fmt.Println("The sum is:", x+y)
}

// Higher Order Function
func processOperation(a int, b int, operation func(x int, y int)) {
	operation(a, b)
}

func call() func(int, int) {
	return add
}

func main() {
	fmt.Println("This is the main function!")
	// fmt.Println(a)

	// anonymous function
	// Immediately Invoked Function Expression (IIFE)
	// func(a int, b int) {
	// 	c := a + b
	// 	fmt.Println("The sum is:", c)
	// }(5, 10)

	// Function expression or Assign a function to a variable
	// add := func(a int, b int) {
	// 	c := a + b
	// 	fmt.Println("The sum is (from function expression):", c)
	// }
	// add(5, 10)

	// processOperation(2, 5, add)
	sum := call()
	sum(3, 4)
}

// func init() {
// 	fmt.Println("This is the init function. It runs before main.")
// 	fmt.Println(a)
// 	a = 20
// }


/*
1. parameter vs argument
2. First Order Functions
	i. standard function / named function
	ii. anonymous function
	iii. Immediately Invoked Function Expression (IIFE)
	iv. function expression or assign a function to a variable
3.Higher Order Functions / First Class Functions
4. Callback Function: A callback function is a function that is passed as an argument to another function and is executed after some operation has been completed. It allows for asynchronous programming and is commonly used in event handling, timers, and asynchronous operations.

functional paradigm -> haskel, racket

** Higher Order Function if one of the following is true:
1. It takes one or more functions as arguments (i.e., procedural parameters).
2. It returns a function as its result (i.e., functional return type).
3. It does both (i.e., it takes one or more functions as arguments and returns a function as its result).
*/