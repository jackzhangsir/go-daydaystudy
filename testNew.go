package main

import (
	"fmt"
	"go/types"
	"reflect"
)

func main() {
	newMap()
	sliceMap()
}

func newMap()  {
	var nmap = new(map[int]string)
	fmt.Println("new的数据类型", reflect.TypeOf(nmap))

	
}
func sliceMap()  {
	var smap = make(map[int]string)
	fmt.Println("make的数据类型", reflect.TypeOf(smap))
}
func newMap2()  {
	var nmap = types.Map{}
	nmap.Key()

}