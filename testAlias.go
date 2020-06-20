package main

import (
	myalias "fmt"
	. "myProject/packx"
	"myProject/packy"
	_ "myProject/packy"
	"time"
)

func main() {

	myalias.Println("hello world")
	myalias.Println(time.Now())
	packy.Show()
	Learn()
	myalias.Println("自己定义的包别名")
	
}
