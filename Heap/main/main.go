package main

import (
	"fmt"
	"github.com/iceriverdog/hellogo/Heap/heap"
)

func main() {
	h := heap.NewHeap()

	arr := []int{8, 1, 4, 5, 2}
	for _, v := range arr {
		h.Insert(v)
	}

	fmt.Println(h.Element)
	h.Delete()
	fmt.Println(h.Element)
	h.Delete()
	fmt.Println(h.Element)
	h.Delete()
	fmt.Println(h.Element)
	h.Delete()
	fmt.Println(h.Element)
	h.Delete()
	fmt.Println(h.Element)
}
