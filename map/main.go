package main

import "fmt"

func main() {
	page := []string{"1", "3"}
	limit := []string{"10"}

	s := map[string] []string{
		"page": page,
		"limit": limit,
	}

	// s := map[string]int{
	// 	"page": 1,
	// 	"limit": 15,
	// }

	// s := map[int]int{
	// 	1: 1,
	// 	2: 5,
	// 	3: 12,
	// }

	val := s["limit"]
	fmt.Println(val)
}
