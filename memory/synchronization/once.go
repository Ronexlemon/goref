package synchronization

import (
	"fmt"
	"sync"
)

var lo sync.Once

var abc string
var tokensStart int

func setup(){
	abc="Game Initilaizing ..."
	tokensStart=10
}

func SendGameTokens(s chan<-int){
	s<-tokensStart
}

func MultiplyGameTokensAfterwin(s chan<-int,r <-chan int){
	t:=<-r
	t=2*t
	s<-t

}

func RunGame(){
	sender:=make(chan int)
	receiver:=make(chan int)

	lo.Do(setup)

	go SendGameTokens(sender)

	go MultiplyGameTokensAfterwin(receiver,sender)
    value := <-receiver

	fmt.Println("The value:", value)





}