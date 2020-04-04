package main

import (
	"fmt"
	"github.com/iceriverdog/hellogo/BinaryTree/binarySearchTree"
)

func main() {
	arr := []int{5, 3, 2, 6, 8, 4}
	tree := binarySearchTree.NewTree(20)
	fmt.Println(len(tree))
	for _, v := range arr {
		tree = tree.Insert(v)
	}
	fmt.Println(tree)
}
