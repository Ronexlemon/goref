package goroutine

import (
	"fmt"
	"time"
)


func Say(word string){
	for i:=0;i<5;i++{
		time.Sleep(time.Second)
		fmt.Println(word)

	}
}

func RunSay(){
	go Say("hello")
	Say("yes")
}