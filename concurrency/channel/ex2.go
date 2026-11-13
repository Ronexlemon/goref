package channel

import (
	"fmt"
	"sync"
)

func Print5Numbers(){
	c:=make(chan int)

	var wg sync.WaitGroup

	for i:=0;i<5;i++{
		wg.Add(1)
		go func (n int)  {
			c<-n
			defer wg.Done()
			
		}(i)
		
	}
 go func ()  {
	wg.Wait()
	close(c)
	
 }()

	for value:=range c{
		fmt.Println("Values",value)
	}
}