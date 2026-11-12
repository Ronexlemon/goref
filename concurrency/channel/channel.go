package channel

import "fmt"


func sum(s []int,c chan int){
	sum:=0
	for _,value:=range s{
		sum+=value
		
	}
	c<-sum
}

func RanSumChannel(){
	s:=[]int{1,2,3,4,5,6,7,8,9,10}
	c:=make(chan int)

	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x,y:=<-c,<-c
	fmt.Printf("The Sums %v and %v sum up to %v",x,y,x+y)

}

func BufferedChannel(){
	c:=make(chan int,2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func OverloadBufferedChannel(){
	c:=make(chan int,2)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	fmt.Println(<-c)
	fmt.Println(<-c)
}