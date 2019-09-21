package main

import (
	"fmt"
	"sync"
	"time"
)

// hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh
func main() {
	list := []int{3, 5, 1, 10, 4, 2}
	wg := sync.WaitGroup{}
	for _, l := range list {
		wg.Add(1)
		go func(value int) {
			time.Sleep(time.Millisecond * 200 * time.Duration(value))
			fmt.Println(value)
			wg.Done()
		}(l)
	}
	wg.Wait()
}
