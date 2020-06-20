package main

import (
	"fmt"
	"time"
)

func main() {

	//这其实就是一个死循环
	One:
		fmt.Println("这其实是一个死循环")
		time.Sleep(1 * time.Second)
	goto One

	//break只终止当前循环
	for i:=0;i< 3;i++{
		for j:=0;j<2;j++{
			fmt.Println("helo")
			break;
		}
	}

	//goto可以跳出多层循环



	
}
