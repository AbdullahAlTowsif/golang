package main

import "fmt"

// func add (num1 int, num2 int) {
// 	sum := num1 + num2
// 	fmt.Println("The sum is:", sum)
// }

// we can also write the above function like this
func add (num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumbers (num1 int, num2 int) (int , int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func myInfo (name string, age int, city string) {
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")
	fmt.Println("I live in", city)
}

func printWelcomeMessage() {
	fmt.Println("Welcome to the application")
}

func getUsername() string {
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var number1 int
	var number2 int
	fmt.Println("Enter first number - ")
	fmt.Scanln(&number1)
	fmt.Println("Enter second number - ")
	fmt.Scanln(&number2)
	return number1, number2
}

func printGoodbyeMessage() {
	name := getUsername()
	fmt.Println("Hello", name)
	fmt.Println("Thank you for using the application. Goodbye!")
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Currently Learning Golang")

	// variable & data types
	a := 10
	var b = 11.43
	var c int = 12
	fmt.Println(a, b, c)

	x := true
	x = false

	fmt.Println(x)

	// if-else | && | ||
	age := 20
	gender := "male"

	if age == 20 && gender == "male" {
		fmt.Println("You are eligible to be married")
	} else if age < 18 {
		fmt.Println("You are not eligible to be married, but you can love someone")
	} else {
		fmt.Println("You are just a teenager, not eligible to be married")
	}

	m := 3

	switch m {
	case 1:
		fmt.Println("m is 1")
	case 2:
		fmt.Println("m is 2")
	case 3:
		fmt.Println("m is 3")
	default:
		fmt.Println("m is not 1, 2, or 3")
	}

	// function call
	result := add(5, 11)
	fmt.Println("The sum is:", result)

	sum, mul := getNumbers(5, 11)
	fmt.Println("The sum is:", sum)
	fmt.Println("The multiplication is:", mul)

	myInfo("John", 25, "New York")

	// Why functions are important?
	
	printWelcomeMessage()

	number1, number2 := getTwoNumbers()
	summation := add(number1, number2)
	fmt.Println("The sum of the two numbers is:", summation)

	// Print goodbye message
	printGoodbyeMessage()
}
