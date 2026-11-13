package synchronization

import (
	"fmt"
)

var a string
func f(){
	fmt.Println(a)
}

func Hello(){
	a ="hello there"
	go f()
	
}