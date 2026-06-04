package main

import "fmt"

var (
	arr2 = [3]string{"Go", "is", "awesome"}
)

func main() {
	// var arr [2]int
	// arr[0] = 1
	// arr[1] = 5

	arr := [2]int{1, 5}
	fmt.Println(arr)
	fmt.Println(arr2)
}
