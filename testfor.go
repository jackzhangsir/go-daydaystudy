package main

import (
	"fmt"
)

func main() {

	var arr = [5]int{1, 22, 33, 44, 55}
	fmt.Println(arr)

	//for{
	//	fmt.Println("你好啊")
	//	//休眠一秒
	//	time.Sleep(1* time.Second)
	//}

	for i := 0; i< len(arr); i++{
		fmt.Println(arr[i])
	}

	aaa := []string{"苹果", "香蕉", "呵呵"}
	//直接遍历值
	for _,v:= range aaa{
		//fmt.Println(k)
		fmt.Println(v)
		break
	}
	//遍历key、value键值对
	for k,v:= range aaa{
		if k == 0{
			continue
		}
		fmt.Println(k)
		fmt.Println(v)
	}

	////for死循环
	//for  {
	//	fmt.Println("hello")
	//}

}
