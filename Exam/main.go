package main

import "fmt"

func IsExist(s byte, m map[uint8]int) bool {
	_, ok := m[s]
	return ok
}

func TypeOfString(s string) int {
	a := make(map[uint8]int)
	if s == "" {
		return 0
	} else if len(s) == 1 {
		return 1
	} else {
		a[s[0]] = 1
		for i := 1; i < len(s); i++ {
			if IsExist(s[i], a) {
				k := 1
				for j := 1; j <= a[s[i]] && i + j < len(s); j++ {
					if s[i] != s[i+j] {
						k = 0
						break
					}
				}
				if k == 1{
					a[s[i]]++
				}
			}
			a[s[i]] = 1
		}
	}
	var total int
	for _, v := range a {
		total += v
	}
	return total
}
func Who(arr []string) {
	for i := 0; i < len(arr); i++ {
		//var m map[uint8]int
		m := make(map[uint8]int)
		s := arr[i]
		for j := 0; j < len(s); j++ {
			if _, ok := m[s[j]]; !ok {
				m[s[j]] = 1
			} else {
				m[s[j]]++
			}
		}
		ji := 0
		for _, v := range m {
			ji += v
		}
		if ji%2 == 1 || len(m) == 1{
			fmt.Println("wk")
		} else {
			fmt.Println("tw")
		}
	}
}

func main() {
	//var s string
	//fmt.Scanln(&s)
	//n := TypeOfString(s)
	//fmt.Println(n)
	var arr []string
	var n int
	fmt.Scanln(&n)
	for i := 1; i <=n; i++ {
		var  s string
		fmt.Scanln(&s)
		arr = append(arr, s)
	}
	fmt.Println(arr)
	//arr := []string{"aba", "ab"}
	Who(arr)
}