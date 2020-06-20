package main

import "fmt"

//定义类型别名
type myint int8

func main() {

	var a myint
	var i int
	var j float32
	var c string
	var d bool

	fmt.Println("myint类型")
	fmt.Println(a)
	fmt.Println("int类型")
	fmt.Println(i)
	fmt.Println("float类型")
	fmt.Println(j)
	fmt.Println("--------------------")
	fmt.Println("string类型")
	fmt.Println(c)
	fmt.Println("bool类型")
	fmt.Println(d)



}
