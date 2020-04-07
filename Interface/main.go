package main

import "fmt"

type Sayer interface{
	say()
}

type dog struct {
	name string
}

type cat struct{
	name string
}

type person struct {
	name string
}

func (d dog) say() {
	fmt.Println("汪汪汪～")
}

func (c *cat) say()  {
	fmt.Println("喵喵喵～")
}

func (p person) say() {
	fmt.Println("啊啊啊～")
}


func main() {
	d := dog{name:"小黑"}
	var s Sayer
	d.say()
	s = d
	s.say()
	c := cat{name:"小白"}
	c.say()
	s = c
	s.say()
}
