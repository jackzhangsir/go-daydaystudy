package main

import (
	"bytes"
	"fmt"
)


var buff bytes.Buffer

func main() {

	arr := []string{"abc", "cde" ,"qwer","uiop"}
	for i:= 0; i < len(arr); i++ {
		buff.WriteString(arr[i]) // 将字符串str写入缓存buffer
	}
	fmt.Println(buff.String())



	
}
