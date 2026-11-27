package main

import (
	"fmt"

	"github.com/ronexlemon/refresh/channel"
)

func main() {
	ch := make(chan channel.Request)
	done := make(chan int64)

	go channel.Counter(ch, done)

	num1 := []int64{1, 2, 3, 4, 5, 5, 6, 7, 7, 8, 9, 3, 2, 1}
	num2 := []int64{2, 3, 4, 5, 6, 7, 89, 3}

	// send inputs
	go func() {
		ch <- channel.Request{Num: num1}
		ch <- channel.Request{Num: num2}
		close(ch)
	}()

	// receive results
	var total int64
	for v := range done {
		total += int64(v)
	}

	fmt.Println("Final sum:", total)
}