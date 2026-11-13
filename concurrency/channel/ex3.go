package channel

import (
	"fmt"
)


func UnbufferedChannel(){
	c:=make(chan int, 3)
	c<-10
	c<-20
	c<-30

	go func(){close(c)}()

	for value:=range c{
		fmt.Println(value)
	}

}