package main

import "fmt"

func main() {

	//var high int = 100
	//switch high {
	//case 80:
	//	fmt.Println("成绩不错")
	//case 90:
	//	fmt.Println("成绩优秀")
	//case 100:
	//	fmt.Println("成绩满分")
	//default:
	//	fmt.Println("以上都不是")
	//}


	var a  interface{}
	a = "nihao"
	switch a.(type) {
	case int:
		fmt.Println("a是整形")
	case string:
		fmt.Println("a是字符串")
	case bool:
		fmt.Println("a是布尔类型")
	case complex64:
		fmt.Println("a是复合类型")
	default:
		fmt.Println("以上都不是")


	}
	
}
