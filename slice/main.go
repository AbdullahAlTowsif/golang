package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

// variadic function to print the slice and its length and capacity
func print(numbers ...int) {
	fmt.Println("From Variadic Function")
	fmt.Println(numbers)
	fmt.Println("length:", len(numbers))
	fmt.Println("capacity:", cap(numbers))
}


func main() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // ptr = 1, len = 3, cap = 5
	fmt.Println(s)

	s1 := s[1:2] // ptr = 2, len = 1, cap = 4
	fmt.Println(s1)

	sl := []int{1, 2, 3, 4, 5} // slice literal
	fmt.Println("slice:", sl, "len:", len(sl), "cap:", cap(sl))

	slice := make([]int, 3) // make a slice with len = 3 and cap = 3
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))

	sliceWithCap := make([]int, 3, 5) // make a slice with len = 3 and cap = 5
	sliceWithCap[0] = 1
	sliceWithCap[2] = 1
	// sliceWithCap[3] = 3 // this will cause a runtime error because the length of the slice is 3, but we are trying to access index 3 which is out of range
	fmt.Println("slice with cap:", sliceWithCap, "len:", len(sliceWithCap), "cap:", cap(sliceWithCap))

	var s2 []int // nil slice
	fmt.Println("nil slice:", s2, "len:", len(s2), "cap:", cap(s2))

	s2 = append(s2, 1) // appending to a nil slice will create a new slice with len = 1 and cap = 1
	fmt.Println("after appending to nil slice:", s2, "len:", len(s2), "cap:", cap(s2))

	// Example to show that slices are reference types and they point to the same underlying array
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println("x:", x) // x: [10 2 3 5]
	fmt.Println("y:", y) // y: [10 2 3 5]

	// Another example
	p := []int{1, 2, 3, 4, 5}
	p = append(p, 6)
	p = append(p, 7)

	q := p[4:]

	r := changeSlice(q) // when passing a slice to a function, it is passed by value, but the underlying array is shared, so changes to the slice inside the function will affect the original slice

	fmt.Println("p:", p) // p: [1 2 3 4 10 6 7]
	fmt.Println("r:", r) // r: [10 6 7 11]
	fmt.Println(p[0:8]) // p: [1 2 3 4 10 6 7 11]

	// calling the variadic function
	print(1, 2, 3, 4, 5)

}

/*
1. slice from an existing array
2. slice from a slice
3. slice literal
4. make function with len
5. make function with len and cap
6. empty slice or nil slice
7. slice underlying array rule => for first 1024 elements, the capacity will be doubled, and for 1025 and above, the capacity will be increased by 25% of the current capacity
*/


/*
First ei slice create hbe na. First e ekta array create hbe. Then oi array te slice create hbe. Then slice er pointer, length and capacity set hbe.

Jodi notun slice append kori tokhon koi ekta jinish check kora lagbe.

Initially append function er jonno ekta stack frame create hbe as well as a new slice. Current slice (mane j slice ta array theke create hobe) er pointer, length and capacity set hbe (more specifically copy hbe) append function er maddhome orthat append function er slice e.

Case 1: Then first of all jodi amader main stack frame er array te jodi **cap > len** hoy tahole notun j value append krbo oita oi array pointer follow kore current array tei append hoye jabe.
But second of all onno case ta hocche.
Eita shudhu main stack frame er jonno na heap memory teo hbe.

Case 2:Jodi main stack frame er array te **cap <= len** hoy tahole GO escape analysis kore completely new ekta array banabe HEAP memory te.
Tokhon amader current j pointer ta ager array te point kore chilo (means main stack frame er array te or HEAP memory er array te)
oi pointer ta change hoye HEAP memory er array te point korbe.

Ei khetre jodi ager array ta HEAP memory te hoy and **cap <= len** hoy tahole GO escape analysis kore abar completely new ekta array banabe HEAP memory te. Tokhon amader current j pointer ta ager array te point kore chilo (means HEAP memory er array te)
oi pointer ta change hoye abar HEAP memory er new array te point korbe.

*/
