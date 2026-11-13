package synchronization

import (
	"fmt"
)

var c = make(chan int,1)

var b string

func ff(){
	b="hello lemon"
	c<-0
}

func SyncThroughChannel(){
	go ff()
	<-c
	fmt.Println(b)
}