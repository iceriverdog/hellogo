package linkedList

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

type Linkedlist struct{
	head *Node
}

func NewNode(value int) *Node {
	return &Node{value:value, next:nil}
}

func (linkedlist *Linkedlist) Insert(value int) *Linkedlist {
		node := NewNode(value)
		temp := linkedlist.head
		linkedlist.head = node
		node.next = temp
		return linkedlist
}

func (linkedlist *Linkedlist) Travel() {
	if linkedlist.head == nil {
		return
	}
	temp := linkedlist.head
	for temp != nil {
		fmt.Print(temp.value)
		temp = temp.next
	}
	return
}

func (Linkedlist *Linkedlist) Length() int {
	temp := Linkedlist.head
	i := 0
	for temp != nil {
		i++
		temp = temp.next
	}
	return i
}

func (linkedlist *Linkedlist) Verse() *Linkedlist {
	if linkedlist.head == nil || linkedlist.head.next == nil {
		return linkedlist
	} else {
		//linked := &Linkedlist{}
		//cur := linkedlist.head
		//fmt.Println(cur)
		//for cur != nil {
		//	linked.Insert(cur.value)
		//	cur = cur.next
		//就地翻转
		cur := linkedlist.head.next
		linkedlist.head.next = nil
		for cur != nil {
			temp := cur.next
			cur.next = linkedlist.head
			linkedlist.head = cur
			cur = temp
		}
		return linkedlist
	}
}

func Verse(node *Node) *Node {
	if node == nil || node.next == nil {
		return node
	}
	newNode := Verse(node.next)
	node.next.next = node
	node.next = nil
	return  newNode
}
func (Linkedlist *Linkedlist) DeleteHead() *Node {
	return  Linkedlist.head
}

func (node *Node) AddHead() Linkedlist {
	return Linkedlist{node}
}

func (linkedlist *Linkedlist) InsertBySorted(value int) *Linkedlist {
	if linkedlist.head == nil {
		linkedlist.head = &Node{value:value}
		return linkedlist
	} else if value <= linkedlist.head.value {
		node := &Node{value:value}
		node.next = linkedlist.head
		linkedlist.head = node
		return linkedlist
	} else {
		temp := linkedlist.head.next
		pre := linkedlist.head
		for temp != nil && value > temp.value {
			pre = temp
			temp = temp.next
		}
		node := &Node{value:value}
		node.next = temp
		pre.next = node
		return linkedlist
	}
}