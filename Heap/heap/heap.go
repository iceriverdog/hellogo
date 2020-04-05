package heap

import (
	"fmt"
	"math"
)

//二叉堆是一种特殊的堆，它满足两个性质：结构性和堆序性
//结构性：二叉堆是一颗完全二叉树，完全二叉树可以用一个数组表示，不需要指针，所以效率更高。
//堆序性质：堆的最小值或最大值在根节点上，所以可以快速找到最大值或最小值。
//当用数组表示时，数组中任一位置i上的元素，其左儿子在位置2i上，右儿子在位置(2i+ 1)上，其父节点在位置(i/2)上。

type Interface interface {
	Len(heap *MaxHeap) int
	swap(i, j int)
	Bigger(i, j int) bool
	Insert(value int)
	Delete()
}

//Define MaxHeap
type MaxHeap struct {
	Element []int
}

// New a Heap
func NewHeap() *MaxHeap {
	//math.MinInt64 --> Integer limit values.
	return &MaxHeap{Element: []int{math.MaxInt64}}
}

// Implementing Methods
func (heap *MaxHeap) Len(i, j int) int {
	return len(heap.Element) - 1
}

func (heap *MaxHeap) Swap(i, j int) {
	heap.Element[i], heap.Element[j] = heap.Element[j], heap.Element[i]
}

func (heap *MaxHeap) Bigger(i, j int) bool {
	return heap.Element[i] > heap.Element[j]
}

func (heap *MaxHeap) Insert(value int) {
	heap.Element = append(heap.Element, value)
	i := len(heap.Element) - 1
	//上浮
	for ; heap.Bigger(i, i/2); i /= 2 {
		heap.Swap(i, i/2)
	}
}

func (heap *MaxHeap) Delete() {
	if len(heap.Element) == 1 {
		fmt.Println("empty heap")
		return
	}
	heap.Element[1] = heap.Element[len(heap.Element)-1]
	heap.Element = heap.Element[:len(heap.Element)-1]

	i := 1
	//Assuming
	MaxPos := i
	// If 2i and 2i + 1 exist , find MaxPos of (heap.Element[i], heap.Element[2i], heap.Element[2i+1])
	// and swap(i, MaxPos)
	for {
		// Find MaxPos
		if 2*i+1 <= len(heap.Element)-1 && heap.Bigger(2*i+1, i) {
			MaxPos = 2*i + 1
		}
		if 2*i <= len(heap.Element)-1 && heap.Element[MaxPos] < heap.Element[2*i] {
			MaxPos = 2 * i
		}
		//判断是否进行交换
		if MaxPos == i {
			return
		} else {
			heap.Swap(i, MaxPos)
			i = MaxPos
		}
	}
}
