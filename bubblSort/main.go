package main

import "fmt"

func bubbleSort(arr *[5]int) {

	temp := 0
	for i := 0; i < len(*arr) - 1; i++ {
		for j := 0; j <len(*arr) - 1 - i; j++{

			if (*arr)[j] > (*arr)[j + 1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}
}
func main() {
	arr := [5]int {2,5,4,3,8}
	bubbleSort(&arr)
	fmt.Println(arr)
}