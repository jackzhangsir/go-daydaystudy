package main

import (
	"fmt"
	"unsafe"
)

//可见性变量
var B_TEST = "hijklmn"
func main() {

	var i uint32 = 9
	fmt.Println(unsafe.Sizeof(i))

	var j uint64 = 8
	fmt.Println(unsafe.Sizeof(j))

	var k int = 88
	fmt.Println(unsafe.Sizeof(k))

	var flag bool = true
	fmt.Println(flag)

	var flag2 bool = false
	fmt.Println(flag2)
	
}
