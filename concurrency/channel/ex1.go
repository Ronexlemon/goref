package channel

import "fmt"


func Print50(){
	c:=make(chan int)

	go func ()  {
		c<-50
		defer close(c)
	}()

	value:=<-c
	fmt.Println("The value",value)
}