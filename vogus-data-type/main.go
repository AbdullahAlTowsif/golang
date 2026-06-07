package main

import "fmt"

func main() {
	var a int8 = -128
	var p int = 67

	var x uint8 = 10 // unsigned(0 and only positive numbers)

	var j float32 = 3.14
	var k float64 = 3.14843

	var flag bool = true // 8 bits --> 1 byte --> %v

	r := '💖' // rune --> alias for int32 (unicode point) --> 32 bits --> 4 bytes --> %c

	var s string = "my name is towsif"
	fmt.Printf("%s\n", s)
	fmt.Printf("%T\n", s)

	fmt.Printf("%c\n", r)
	fmt.Println(a, p, x, j, k, flag)
}
