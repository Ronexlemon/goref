package channel

import (
	"fmt"
	"sync"
	"time"
)


type Safecounter struct{
	mu sync.Mutex
	v map[string]int
}

func(s *Safecounter)Increment(key string){
	s.mu.Lock()
	s.v[key]++
	s.mu.Unlock()
}


func (s *Safecounter)Value(key string)int{
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func MutualExclusion(){
	key:="someKey"
	s:=Safecounter{v:map[string]int{}}

	for i:=0;i<1000;i++{
		go s.Increment(key)
	}
	time.Sleep(time.Second)
	fmt.Println("The key",s.Value(key))
}