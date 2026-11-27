package synclock

import "sync"

type Counter struct{
	mu sync.Mutex
	sum int64
}

var mu sync.Mutex




func(c *Counter) Add(num []int64){
	mu.Lock()
	defer mu.Unlock()
	for _,v:=range num{
		c.sum += v
	}
	
}

func (c *Counter) Value()(int64){
	mu.Lock()
	defer mu.Unlock()
	return c.sum
}