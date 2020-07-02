package main

import "fmt"

func main() {

	var intArr = []int{11, 22, 33, 44, 555}
	//打印长度和容量
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
	//向数组中追加元素
	intArr = append(intArr, 55)
	//for-range循环
	for _,v :=  range intArr{
		fmt.Println(v)
	}

	//
	for i:= 0;i<len(intArr);i++{
		fmt.Println(intArr[i])
	}
	//
	fmt.Println("长度为", len(intArr))
	fmt.Println("容量为", cap(intArr))
	
}
