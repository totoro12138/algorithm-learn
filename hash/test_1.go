package main

import (
	"fmt"
	"hash/crc32"
	"log"
)

func main() {
	input := "dfewrwed"
	// 将字符串转换成唯一的hashcode
	hashCode := crc32.ChecksumIEEE([]byte(input))
	fmt.Printf("hashcode: %v\n", hashCode)
}

type HashMap struct {
	List   []*Node
	Length int
}

type Node struct {
	Key   string
	Value string
	Next  *Node
}

func NewMap() HashMap {
	return HashMap{
		List:   make([]*Node, 10),
		Length: 10,
	}
}

func (h *HashMap) get(key string) (value string) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	var p *Node
	index := int(crc32.ChecksumIEEE([]byte(key))) % h.Length
	if h.List[index] != nil {
		p = h.List[index]
	}

	return ""
}

func (h *HashMap) put(key, value string) {

}

func (h *HashMap) del() {

}
