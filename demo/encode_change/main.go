package main

import (
	"fmt"
)

func main() {
	var dst []byte
	var dst2 []rune
	src := "fads1256发的架飞12312fdsa"
	dst = []byte(src)
	dst2 = []rune(src)

	//ret := strconv.AppendQuote(dst,src)
	//fmt.Println("src:",[]byte(src))
	//fmt.Println("src str:",src)
	fmt.Println("dst:", dst)
	fmt.Println("dst str:", string(dst))
	fmt.Println("dst len:", len(dst))
	fmt.Println("dst2:", dst2)
	fmt.Println("dst2 str:", string(dst2))
	fmt.Println("dst2 len:", len(dst2))
	//strconv.QuoteRuneToASCII()
}
