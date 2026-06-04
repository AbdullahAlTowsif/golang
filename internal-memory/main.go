package main

import "fmt"

const p = 100

var a = 10

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println("The sum is:", z)
	}

	add(5, 4)
	add(p, a)
}

// func add(x int, y int) {
// 	z := x + y
// 	fmt.Println("The sum is:", z)
// }

func main() {
	// add(5, 4)
	// add(a, 3)

	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("Hello")
}


/*

2 Phases:
1. compilation phase
2. execution phase

**** Compilation Phase ****

### Code Segment ###

p = 100
call = func() {...}
add = func() {...}
main = func() {...}
init = func() {...}



go run main.go => compile it => main => ./main
go build main.go => compile it => main
./main => execute it

*/