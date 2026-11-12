package channel

import (
	"fmt"
	"time"
)

func fibbanociselect(c,quit chan int){
	x,y:=0,1
	
	for{
		select{
	case c<-x:
		x,y=y,x+y
	case <-quit:
		fmt.Println("quit")
		return
	}
	}
}

func SelectFibboci(){
	c:=make(chan int)
	quit:=make(chan int)

	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-c)

		}
		quit <-0

	}()

	fibbanociselect(c,quit)
}

func TimeSequenceDefault(){
	start:=time.Now()
	tick:=time.Tick(100 *time.Millisecond)
	boom:= time.After(time.Millisecond *1000)

	elapsed:=func()time.Duration{
		return time.Since(start).Round(time.Millisecond)

	}

	for{
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}