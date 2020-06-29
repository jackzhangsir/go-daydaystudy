package sugar_demo

import "fmt"

//
func Sugar(values...string){
	for _,v := range values{
		fmt.Println("v--> ", v)
	}
}

//
func TypeJuge(){
	value := false
	fmt.Println("value = ", value)
}
