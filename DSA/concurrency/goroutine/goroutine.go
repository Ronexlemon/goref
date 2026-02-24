package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func ForeverCount(thing string){
	for i:=0;true;i++{
		fmt.Printf("Index  %d : %s \n",i,thing)
		time.Sleep(time.Second *2)
	}
}

func CountwithChannel(thing string, c chan string){
	for i:=0;i <10;i++{
		c <- thing
		time.Sleep(time.Millisecond *500)
	}
	close(c)
}

func DoCount(){
	go ForeverCount("golang")
	 ForeverCount("Rust")
}
func DocountWithWait(){
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		ForeverCount("compact")
		wg.Done()

	}()
	wg.Wait()
}

func DoCountWithChannel(){
	c := make(chan string)

	go CountwithChannel("Python",c)

	for msg:= range c{
		fmt.Println(msg)
	}
}