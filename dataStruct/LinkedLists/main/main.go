package main

import (
	"fmt"
	"dataStruct/LinkedLists/linkedList"
)

func main() {
	linked := &linkedList.Linkedlist{}
	linked.Insert(1)
	linked.Insert(2)
	linked.Insert(3)
	linked.Insert(4)
	linked.Travel()
	length := linked.Length()
	fmt.Print(length)
	fmt.Println()
	linked2 := linked.Verse()
	linked2.Travel()
	linked3 := linkedList.Verse(linked.DeleteHead())
	linked4 := linked3.AddHead()
	linked4.Travel()
	l := &linkedList.Linkedlist{}
	fmt.Println("从小到大插")
	arr := []int{5,2,3,4,1}
	for _, v := range arr {
		l = l.InsertBySorted(v)
	}
	l.Travel()

}
