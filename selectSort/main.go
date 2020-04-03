package main

import (
	"fmt"
)

func selectSort(arr *[5]int) {

	for j := 0; j < len(arr) - 1; j++ {
		max := arr[j]
		index := j

		for i := j +1; i < len(arr); i++ {

			if max < arr[i] {
				max = arr[i]
				index = i
			}
		}
		if index != j {
			temp:=arr[j]
			arr[j] = arr[index]
			arr[index] = temp
			fmt.Println(arr)
		}
	}
}

func main() {
 
	a := [5]int {10, 20,50,66,33}
	selectSort(&a)           
	fmt.Println(a)
}
