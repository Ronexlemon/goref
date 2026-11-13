package synchronization

import (
	"fmt"
	"time"
)

func doWork(num int ){
	for i:=0;i <10;i++{
		time.Sleep(1*time.Second)
		fmt.Println("The nume:",num)
	}
}

var limit=make(chan int,3)

func RunWork(){

	s:=make([]int,10)
	for i:=0;i<10;i++{
		s[i]=i
	}
	for _,value:=range s{
		go func (w func(nums int))  {
			limit <-1
			doWork(value)
			<-limit
	
		}(doWork)

	}
	select{}

}