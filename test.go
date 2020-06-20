package main

//导入依赖包
import "fmt"


//声明常亮和变量
const NAME = "HAHAHA"
var name string = "jack"
//自定义类型
type myInt int
//声明结构体，相当于类
type MyStruct struct {
	name string
	sex string
	age int
}
//声明接口
type IDemo interface {

}

func myfunc() {
	fmt.Println("hello world from myfunc")
}

//主函数
func main() {
	myfunc()
	fmt.Println("hello world")
	
}
