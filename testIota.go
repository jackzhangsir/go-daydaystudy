package main

import "fmt"

const(
	aaa = iota
	_
	bbb = iota
	_
	ccc = iota
	ddd = 3.1415
	eee = iota * 5

)

func main() {

	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	fmt.Println(eee)
	
}
