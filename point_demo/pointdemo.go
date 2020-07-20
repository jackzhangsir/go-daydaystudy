package point_demo

import "fmt"

func PointTest()  {
	var count int = 20
	var countPoint *int
	countPoint = &count
	fmt.Println("count =", count)
	fmt.Println("countPoint = ",countPoint)
	
}

func PointTest2(){
	//指针数组
	a, b := 1, 2
	pointArr := [...]*int{&a , &b}
	fmt.Println("指针数据=", pointArr)

	//数组指针

	var arrPoint *[]int
	arrPoint = &[]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrPoint)



}

