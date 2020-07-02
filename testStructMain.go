package main

import (
	"fmt"
	"myProject/struct_demo"
)

func main() {


	//struct_demo.TestForStruct()

	var dog = new(struct_demo.Dog)
	dog.ID = 3
	dog.Name ="mike"
	dog.Age = 40
	//dog.printName()
	fmt.Println(dog)
	
}
