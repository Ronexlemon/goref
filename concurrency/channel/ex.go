package channel

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int){
	if t ==nil{
		return
	}
	Walk(t.Left,ch)

	ch<-t.Value
	Walk(t.Right,ch)
}


func Same(t1, t2 *tree.Tree) bool{
	ch1:=make(chan int)
	ch2:=make(chan int)

	go func ()  {
		Walk(t1,ch1)
		close(ch1)
		
	}()
	go func ()  {
		Walk(t2,ch2)
		close(ch2)
		
	}()

	for{
		v1,ok1:=<-ch1
		v2,ok2:=<-ch2

		if ok1 !=ok2{
			return false
		}

		if !ok1 && !ok2{
			return true
		}
		if v1 !=v2{
			return false
		}
	}
}

func RUNEx(){
	ch := make(chan int)

	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
