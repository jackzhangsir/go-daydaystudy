package point_demo

import "fmt"

func PointTest()  {
	var count int = 20
	var countPoint *int
	countPoint = &count
	fmt.Println("count =", count)
	fmt.Println("countPoint = ",countPoint)
	
}
