package main

import (
	//"github.com/ronexlemon/goref/DSA/concurrency/goroutine"
	"github.com/ronexlemon/goref/DSA/concurrency/selectgo"
)


func main(){
	//goroutine.DoCount()
	//goroutine.DocountWithWait()
	//goroutine.DoCountWithChannel()
	selectgo.SelectWithChannel()

}