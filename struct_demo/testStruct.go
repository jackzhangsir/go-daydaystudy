package struct_demo

import "fmt"

//定义一个结构体
type Dog struct {
	ID int
	Name string
	Age int
}

//
func TestForStruct()  {
	//方式一
	//var dog Dog
	//dog.ID = 1
	//dog.Name = "jack"
	//dog.Age = 15

	//方式二
	//dog := Dog{ID: 1,Name: "rose",Age: 13}

	//方式三
	//new 返回的指针
	dog := new(Dog)
	dog.ID = 3
	dog.Name ="mike"
	dog.Age = 40
	fmt.Println("dog = ", dog)
}

//为结构体添加方法
func (d *Dog) PrintName()  {
	fmt.Println("它的名字为", d.Name, "打印完成")
}