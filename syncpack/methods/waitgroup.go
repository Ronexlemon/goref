package methods

import "sync"

func WaigGroup()int{
	mu:=sync.Mutex{}
	wg:=sync.WaitGroup{}
	counter:=0

	for i:=0;i<1000;i++{
		wg.Add(1)
		go func (n int)  {
			defer wg.Done()
			mu.Lock()
			
			counter+=n
			mu.Unlock()
			
		}(i)
	}
	wg.Wait()

	return counter


}

//using atomic 

