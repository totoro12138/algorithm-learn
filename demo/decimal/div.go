package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	price := int64(333333333)
	mod := int64(111111111)
	fmt.Println(decimal.New(price, -2).Mod(decimal.New(mod, -2)).String())
}
