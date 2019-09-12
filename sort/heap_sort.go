package main

import "fmt"

func main() {
	list := Heap{
		7, 1, 3, 10, 5, 2, 8, 9, 6,
	}
	list.buildHeap()
	fmt.Println(list)
}

type Heap []int

func (h Heap) buildHeap() {
	hLength := len(h)
	for i := (len(h) - 1) / 2; i >= 0; i-- {
		h.downAdjust(i, hLength)
	}
}

func (h Heap) upAdjust() {

}

func (h Heap) downAdjust(parentIndex, length int) {
	parentValue := h[parentIndex]
	childIndex := 2*parentIndex + 1
	hLength := len(h)
	for childIndex < hLength {
		// 判断 右结点 是否 小于 左结点
		if childIndex+1 < hLength && h[childIndex+1] < h[childIndex] {
			childIndex++
		}
		// 父结点小于两个子结点，直接退出
		if parentValue < h[childIndex] {
			break
		}

		h[parentIndex] = h[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	h[parentIndex] = parentValue
}
