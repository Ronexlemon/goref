package methods

import (
	"fmt"
	"sync"
)


var numbers []int
var mu sync.RWMutex

func store(x int){
	mu.Lock()
	numbers=append(numbers, x)
	mu.Unlock()

}

func avg()float64{

	mu.RLock()
	defer mu.RUnlock()

	size:=len(numbers)
	sum:=0

	for _,val:=range numbers{
		sum+=val
	}

	return float64(sum)/float64(size)
}

func RWMutex(){
	wg:=sync.WaitGroup{}
	ch:=make(chan int,1000)

	for i:=0;i<1000;i++{
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			store(n)
			
		}(i)
	}


	for i:=0;i<1000;i++{
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			ch <-int(avg())
			
			
		}(i)
	}

	wg.Wait()
	close(ch)

	for n:=range ch{
		fmt.Println(n)
	}
	
}