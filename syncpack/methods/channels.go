package methods

import "sync"



func Channels()int{
	counter:=0

	wg:=sync.WaitGroup{}

	ch:=make(chan int,1000)

	for i:=0;i<1000;i++{
		wg.Add(1)
		go func (n int)  {
			defer wg.Done()
			ch<-i
			
		}(i)
	}
	wg.Wait()
	close(ch)
	for n:=range ch{
		counter+=n
	}
	return counter
}