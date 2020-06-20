package main

import (
	"fmt"
)

func main() {

	i:= 88
	fmt.Println(i)

	j := byte(4)
	fmt.Println(j)

	k := int32(86)
	fmt.Println(k)

	a := float32(5.21)
	fmt.Println(a)

	b := float64(4.01)
	fmt.Println(b)

	arr := [5]int8{11 ,22 ,33 ,44 ,55}
	for i2 := range arr {
		fmt.Println(arr[i2])
	}
	//fmt.Println(len(arr))

}
