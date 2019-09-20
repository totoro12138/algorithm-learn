package main

import "fmt"

//go:nosplit
func Swap(a, b int) (int, int)

func main() {
	fmt.Println(Swap(1, 2))
}
