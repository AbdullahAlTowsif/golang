package main

import "fmt"

func print(numbers *[3]int) {
	fmt.Println("Array inside print function:", numbers)
}

type User struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	// fmt.Println("Hello, Pointer!")

	// Pointer or address of memory (RAM)
	x := 10
	address := &x // & is the address operator
	*address = 20 // changing the value at the address

	fmt.Println("Address of x:", address)
	fmt.Println("Value at the address:", *address) // * is the dereference operator
	fmt.Println("Value of x:", x)

	arr := [3]int{1, 2, 3}
	print(&arr) // passing the array by reference (address of the array)

	obj := User{
		Name:"John Doe",
		Age: 30,
		Salary: 50000.0,
	}

	p := &obj // p is a pointer to obj
	fmt.Println("Pointer to obj:", *p)
	fmt.Println("Name:",p.Name)

	fmt.Println(obj)

}

/*

& --> Before a variable --> Meaning Address-of --> Where is this located in memory?
* --> Before a variable --> Meaning Dereference --> What is the value at this address in memory?
* --> Before a type --> Meaning Pointer Type --> This variable is a pointer to this type


*******More Preciesely*******
& ---> address of
* ---> value at address
*/
