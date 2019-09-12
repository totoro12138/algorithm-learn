package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//list := List{
	//	Values: []int{
	//		16353, 19804, 811, 6738, 7172, 368, 11868,
	//		14771, 10293, 4849, 12414, 5336, 811, 8340,
	//		1669, 5331, 2028, 7390, 9480, 10926, 12464,
	//		7600, 14195, 11773, 12452, 15607, 7895, 16039,
	//		12494, 2606, 2442, 16746, 11265, 10716, 1783,
	//		8258, 3652, 15179, 18517, 9340, 17815, 11040,
	//		12702, 2620, 3465, 2550, 17564, 4015, 19456,
	//		5389, 18057, 12753, 19854, 17331, 11391, 11792,
	//		8485, 14310, 15583, 18738, 7326, 19650, 8002,
	//		1848, 15546, 1586, 15881, 5813, 9617, 13297, 173,
	//		4426, 6090, 5690, 453, 15111, 16235, 14389, 15566,
	//		6958, 15369, 7112, 13577, 19888, 5834, 10056,
	//		1541, 5723, 16457, 15204, 14368, 14815, 14785,
	//		9986, 114, 8248, 566, 13140, 17039, 19045,
	//	},
	//}
	list := List{
		4, 7, 6, 5, 3, 2, 8, 1,
	}
	list.quickSort(0, len(list)-1)
	fmt.Println(list)
}

type List []int

// 双边循环法
func (l List) partitionV1(startIdx, endIdx int) int {
	rand.Seed(time.Now().Unix())
	idx := rand.Int()%(endIdx-startIdx+1) + startIdx
	pivot := l[idx]
	//fmt.Println("[pivot]", pivot)
	fmt.Println("[1]", l[startIdx:endIdx+1])
	l[startIdx], l[idx] = l[idx], l[startIdx]
	fmt.Println("[2]", l[startIdx:endIdx+1])
	left := startIdx
	right := endIdx
	fmt.Println("left", left, "right", right)
	for left != right {
		for left < right && l[right] > pivot {
			right--
		}
		for left < right && l[left] <= pivot {
			left++
		}

		if left < right {
			l[left], l[right] = l[right], l[left]
		}
	}

	l[startIdx], l[left] = l[left], l[startIdx]
	fmt.Println("[3]", l[startIdx:endIdx+1])
	return left
}

// 单边循环法
func (l List) partitionV2(startIdx, endIdx int) int {
	rand.Seed(time.Now().Unix())
	idx := rand.Int()%(endIdx-startIdx+1) + startIdx
	pivot := l[idx]
	//fmt.Println("[pivot]", pivot)
	fmt.Println("[1]", l[startIdx:endIdx+1])
	l[startIdx], l[idx] = l[idx], l[startIdx]
	fmt.Println("[2]", l[startIdx:endIdx+1])
	mark := startIdx
	for i := startIdx + 1; i <= endIdx; i++ {
		if l[i] < pivot {
			mark++
			l[i], l[mark] = l[mark], l[i]
		}
	}

	l[startIdx], l[mark] = l[mark], l[startIdx]
	fmt.Println("[3]", l[startIdx:endIdx+1])
	return mark
}

func (l List) quickSort(startIdx, endIdx int) {
	if startIdx >= endIdx {
		return
	}

	//pivotIdx := l.partitionV2(startIdx, endIdx)
	pivotIdx := l.partitionV1(startIdx, endIdx)
	l.quickSort(startIdx, pivotIdx-1)
	l.quickSort(pivotIdx+1, endIdx)
}
