package channel_demo

import (
	"fmt"
	"time"
)

var chanInt chan int = make(chan int, 10)
var chanbool chan bool = make(chan bool)

func Send()  {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <-2
	time.Sleep(time.Second * 1)
	chanInt <-3

	time.Sleep(time.Second * 2)
	chanbool <- false
}

func Receive()  {

	for  {
		select {
		case num :=<-chanInt:
			fmt.Println("chanInt = ", num)
		case <-chanbool:
			fmt.Println("chanbool = ","timeout" )
		}
	}

}