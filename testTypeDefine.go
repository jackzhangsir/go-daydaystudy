package main

import (
	"fmt"
	"reflect"
)
//全局变量必须用var
var a int
var b string = "456"
//变量分组声明方式
var(
	c string = "abc"
	d int = 66
)
func main() {
	//局部变量可以省略var
	var m, n ,l int = 5, 6, 7
	//第二种赋值方式
	x,y,z := 77, 3.3, "abc"
	//下划线是放弃一个值
	var q, _, e = 12, 34, 45

	fmt.Println(a)
	fmt.Print(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("------------")
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(l)
	fmt.Println(reflect.TypeOf(m))
	fmt.Println("------------")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("------------")
	fmt.Println(q)
	fmt.Println()
	fmt.Println(e)



}
