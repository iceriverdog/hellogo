package main

import (
	"fmt"
	"github.com/iceriverdog/hellogo/Queue/queue"
)
func main() {
	queue := queue.NewQueue(2)
	queue.EnQueue("wk")
	queue.EnQueue("cc")
	err := queue.EnQueue("tw")
	fmt.Println(err)
	queue.Travel()
}
