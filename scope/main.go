package main

import (
	"fmt"
	"anything.com/mathlib"
)

var (
	a = 20
	b = 30
)

func printNum(num int) {
	fmt.Println("The number is:", num)
}

func add(x int, y int) {
	z := x + y
	printNum(z)
}

func main() {
	// p := 30
	// q := 40

	// add(p, q)
	// add(a, b)
	// add(a, p)

	// add(a, z) // This will cause an error because z is not defined in this scope

	/* 🖇️ go mod init anything.com 🖇️ */
	fmt.Println("Showing Custom Package")
	
	// To import func, variable whatever from another package, the func must be exported (i.e., start with an uppercase letter)
	mathlib.Add(10, 20)
	fmt.Println("Money in mathlib package is:", mathlib.Money)

	add(a, b)
}

/*
Types of Scope in Golang:

1. Package Scope: Variables declared at the package level are accessible throughout the entire package. They can be accessed by any function within the same package.
2. Function Scope: Variables declared within a function are only accessible within that function. They cannot be accessed outside of the function.
3. Block Scope: Variables declared within a block (e.g., within an if statement, for loop, or switch case) are only accessible within that block. They cannot be accessed outside of the block.
4. Global Scope: Variables declared outside of any function or block are considered global and can be accessed from anywhere in the program, including other packages if they are exported (i.e., start with an uppercase letter).
5. Local Scope: Variables declared within a function or block are considered local to that function or block and cannot be accessed outside of it.

*/
