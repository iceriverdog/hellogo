package queue

import (
	"errors"
	"fmt"
)

type Node struct {
	value string
	next *Node
}
type Queue struct{
	Max int
	Last int
	First *Node
}

func NewQueue(max int) *Queue {
	return &Queue{max, 0, nil}
}

func (queue Queue) IsEmpty() bool {
	if queue.First == nil {
		return true
	}
	return false
}

func (queue Queue) Travel() {
	if queue.IsEmpty() {
		return
	}
	temp := queue.First
	for temp != nil {
		fmt.Print(temp.value)
		temp = temp.next
	}
}

func (queue *Queue) EnQueue(value string) (err error) {
	if queue.Last < queue.Max {
		if queue.First == nil {
			queue.Last ++
			queue.First = &Node{value:value,next:nil}
			return
		} else {
			temp := queue.First
			for temp.next != nil {
				temp = temp.next
			}
			queue.Last++
			temp.next = &Node{value: value, next: nil}
			return
		}
	}
	return errors.New("queue is full")
}

func (queue Queue) DeQueue() (err error) {
	if queue.IsEmpty() {
		return errors.New("empty queue")
	}
	queue.Last--
	queue.First = queue.First.next
	return
}
