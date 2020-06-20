package main

import (
	"fmt"
	"myProject/packx"
	"myProject/packy"
)

//注意导包的顺序
import (
	"fmt"
	"myProject/packx"
	"myProject/packy"
)

func init(){
	fmt.Println("在main方法之前，先执行")
}

func main() {
	fmt.Println("hello world from godemo")
	//注意包的先后初始化，
	//导入包之后，会先初始化常量和变量，然后执行init初始化函数
	//然后执行当前文件中的init函数，最后执行main函数
	packy.Show()
	packx.Learn()
}
