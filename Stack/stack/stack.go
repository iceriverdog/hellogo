package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Max int
	Top int
	Array []int
}

func NewStack(Max int) *Stack {
	arr := []int{}
	return &Stack{Max:Max, Top:-1, Array:arr}
}

func (stack Stack) IsEmpty() bool {
	if stack.Top == -1 {
		return true
	}
	return false
}

func (stack Stack) Travel() {
		if stack.IsEmpty() {
			return
		} else {
			for i := stack.Top; i > -1; i-- {
				fmt.Print(stack.Array[i])
			}
		}
}

func (stack *Stack) Push(value int) (err error) {
	if stack.Top == stack.Max - 1 {
		return errors.New("stack is full")
	} else {
		stack.Array = append(stack.Array, value)
		stack.Top ++
	}
	return
}
func (stack *Stack) Pop() (err error) {
	if stack.Top == -1 {
		return errors.New("empty stack")
	} else {
		stack.Array = stack.Array[:len(stack.Array) - 1]
		stack.Top --
	}
	return
}

