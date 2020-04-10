package main

import "fmt"

func main() {
	//	链接：https://juejin.im/post/5d12db30e51d455a68490ba7
	// 官方源代码这样定义byte
	// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
	// used, by convention, to distinguish byte values from 8-bit unsigned
	// integer values.

	// Convert between byte array/slice and string

	//b := []byte("ABC€")
	//fmt.Println(b) // [65 66 67 226 130 172]

	//s := string([]byte{65, 66, 67, 226, 130, 172})
	//fmt.Println(s) // ABC€

	//type byte = uint8
	var c1 uint8 = 'a'
	var c2 byte = '0' //字符的0

	//当我们直接输出byte值，就是输出了的对应的字符的码值
	// 'a' ==>
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	// byte 放不下 中文字
	//var c3 byte = '北' //overflpackage main

	var c3 int = '北' //overflow溢出
	fmt.Printf("c3=%c c3对应码值=%d\n", c3, c3)
	fmt.Println(c3)
	// %q 字符对应的字面值，用单引号表示出来
	fmt.Printf("%q", c3)
}
