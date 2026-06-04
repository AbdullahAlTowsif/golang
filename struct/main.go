package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUserDetails(user User) {
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
}

// receiver function
func (user User) printDetails() {
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
}

func (user User) call(a int) {
	fmt.Println(user.Name)
	fmt.Println(a)
}

func main() {
	var user1 User

	user1 = User{
		Name: "Alice",
		Age:  30,
	}

	// fmt.Println("Name:", user1.Name)
	// fmt.Println("Age:", user1.Age)
	// printUserDetails(user1)
	user1.printDetails()
	user1.call(10)

	user2 := User{
		Name: "Bob",
		Age:  25,
	}

	printUserDetails(user2)
}
