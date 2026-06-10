package main

import "fmt"

func a() {
	i := 0
	fmt.Println("first:", i)
	defer fmt.Println("second:", i)

	i++
	fmt.Println("third:", i)
	defer fmt.Println("fourth:", i)
	// return
}

func sum(a int, b int) (result int) {
	result = a + b
	return
}


// defer uses defer list stack pointer
func calculate() (result int) {
	fmt.Println("Inside named returned value function")
	fmt.Println("first:", result)
	show := func() {
		result = result + 10
		fmt.Println("defer:", result)
	}

	// closure --> show(), result --> jodi parent stack frame exist kore tahole reference use hobe, jodi na kore tahole copy of result show() er stack frame e store hobe
	defer show()

	result = 5
	p := func (a int) {
		fmt.Println("ami", a)
	}
	defer p(result)
	defer fmt.Println(result)

	fmt.Println("second:", result)

	defer fmt.Println(5)
	return
}


func calc() int {
	fmt.Println("Inside normal function")
	result := 0
	fmt.Println("first:", result)
	show := func() {
		result = result + 10
		fmt.Println("defer:", result)
	}

	defer show()

	result = 5
	fmt.Println("second:", result)
	return result
}

/*
named return value function: calculate()
1. all codes execute
2. defer function store kore rakhe magic box e
3. return -> all defer functions execute hobe
4. return korbe named variables gular values

just return types
1. all codes execute
2. defer function store kore rakhe magic box e
3. return values are evaluated at this time (store the return value)
4. all defer functions execute hobe
*/

func main() {
	// a()
	// res := sum(1, 2)
	// fmt.Println("result:", res)
	a := calculate()
	// b := calc()
	fmt.Println("===============")
	fmt.Println("Calculate First", a)
	// fmt.Println("Calc Second", b)
}

// 1:19:00