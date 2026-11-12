package channel

import "fmt"

func fibbanoci(n int, c chan int){
	x,y:=0,1
	for i :=0; i< n; i++{
		c <-x
		x,y=y,x+y
	}
	close(c)
}

func FindFibbanoci(){
	n:=100
	ch:=make(chan int)
	go fibbanoci(n,ch)
	for i:=range ch{
		fmt.Print(i)
	}
}