package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)

	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func main() {
	fmt.Println(string(Krand(5, 0)))
}
