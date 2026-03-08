package exer

import (
	"fmt"
	"time"
)


func Worker(id int,jobs <-chan int, results chan <-int){
	for j := range jobs{
		fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
	}

}

func WorkerRuns(){
	numbOfJobs:=5

	jobs := make(chan int,numbOfJobs)
	results := make(chan int,numbOfJobs)

	for w:=1; w<=3;w++{
		go Worker(w,jobs,results)

	}

	for j:=1; j<=numbOfJobs;j++{
		jobs <- j
	}
	close(jobs)

	for r:=1;r <=numbOfJobs;r++{
		fmt.Println(<-results)
	}
}