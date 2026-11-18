package main

import (
	"fmt"

	"github.com/ronexlemon/dsa/syncypack/methods"
)



func main(){
	print(methods.WaigGroup())

}

func print(args interface{}){
	fmt.Println(args)
}