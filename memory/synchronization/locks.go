package synchronization

import (
	"fmt"
	"sync"
)


var l sync.Mutex

var ab string

func fff(){
	ab="acquired the lock about to unlock"
	l.Unlock()
}

func RunLock(){
	l.Lock()
	go fff()
	l.Lock()
	fmt.Println(ab)
}

//rmultiple goroutines but one writer
var mu sync.RWMutex
var balance int

func readerBalance(){
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("The balance",balance)
}

func updateBalance(){
	mu.Lock()
	defer mu.Unlock()
	balance+=100
	
}

func RunRWLocks(){
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			updateBalance()
			
		}()
	
	}

	for i:=0;i<10;i++{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			readerBalance()
			
		}()
	
	}
	wg.Wait()


	fmt.Println("Final balance:", balance)
}