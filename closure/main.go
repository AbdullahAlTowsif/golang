package main

import "fmt"

const a = 10
var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age =", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("init() called")
}


/*
2 phases:
1. compilation phase (compile time)
2. execution phase (run time)


***********Compilation phase***********

***Code Segment***

a = 10
outer = func() {...}
outerAnonymous1 = func() {...}
call = func() {...}
main = func() {...}
init = func() {...}

***********Execution phase***********
---> First it checks from line1 to finish line and stores the code in RAM and then it starts executing the code from main() function. But before executing main() function, it executes init() function.


*/
