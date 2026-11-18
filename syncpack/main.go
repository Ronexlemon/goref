package main

import (
	"fmt"

	"github.com/ronexlemon/dsa/syncypack/methods"
)



func main(){
	// print(methods.WaigGroup())
	// print(methods.Atomic())
	// print(methods.Channels())
	methods.RWMutex()

}

func print(args interface{}){
	fmt.Println(args)
}