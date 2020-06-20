package main

import "fmt"

func main() {
	//makeSlice()
	makeMap()
}
//
func makeSlice()  {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"
	fmt.Print(mSlice)
}
//
func makeMap()  {
	mmap := make(map[int]string)
	mmap[0] = "dog"
	mmap[1] = "cat"
	fmt.Print(mmap)
}