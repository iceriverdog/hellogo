package main

import (
	"fmt"
	"github.com/iceriverdog/hellogo/Stack/stack"
)

func main() {
	//New stack
	stack := stack.NewStack(5)
	//Is empty ?
	boolean := stack.IsEmpty()
	fmt.Println(boolean)
	//Push
	arr := [6]int{1, 2, 3, 4, 5, 6}
	for _, v := range arr {
		err := stack.Push(v)
		if err == nil {
			fmt.Printf("push %d succeed", v)
			fmt.Println()
		} else {
			fmt.Println(err)
		}
	}
	stack.Travel()
	//Pop
	stack.Pop()
	stack.Travel()

}
