package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := ""
	fmt.Println(atoi(a))

}

func atoi(value string) (int, error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return strconv.Atoi(value)
}
