package main

import (
	"fmt"
	"sync"
)

var cnt int64
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			
			a := cnt
			a = a + 1
			cnt = a
			
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}