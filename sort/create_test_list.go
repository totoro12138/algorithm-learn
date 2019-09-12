package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		fmt.Printf("%v,", rand.Int()%20000)
		time.Sleep(time.Millisecond * 100)
	}
}
