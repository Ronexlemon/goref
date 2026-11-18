package methods

import (
	"sync"
	"sync/atomic"
)


func Atomic()int64{
	wg:=sync.WaitGroup{}
	var counter  atomic.Int64

	for i:=0;i<1000;i++{
		wg.Add(1)
		go func (n int64)  {
			defer wg.Done()
			counter.Add(n)

			
		}(int64(i))
	}
 wg.Wait()
	return counter.Load()
}