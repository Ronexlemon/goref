package selectgo

import (
	"fmt"
	"time"
)


func SelectWithChannel(){
	c1:=make(chan string)
	c2:=make(chan string)

	go func ()  {
		for{
			time.Sleep(time.Millisecond *500)
			c1 <- "channel 1"
		}
		
	}()
	go func ()  {
		for{
			time.Sleep(time.Second *2)
			c2 <- "channel 2"
		}
		
	}()

	for{
		select{
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		
		}
	}
}
