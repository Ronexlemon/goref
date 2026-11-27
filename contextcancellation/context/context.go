package context

import (
	"context"
	"fmt"
	"net/http"
	"time"
)


func RunContext(){
	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()

	req,err:=http.NewRequestWithContext(ctx,"GET","https://example.com",nil)
	if err !=nil{
		fmt.Println("Failed to get the request")
		return 
	}
	client:=&http.Client{}
	res,err:=client.Do(req)
	if err !=nil{
		fmt.Println("The error",err)
		return 
	}

	defer res.Body.Close()
	fmt.Println("The status",res.Status)
}

func task(ctx context.Context,id int){

	select{
	case <-time.After(time.Duration(id)*time.Second):
		 fmt.Printf("Task %d done!\n", id)
	case <-ctx.Done():
		fmt.Printf("Task %d bailed: %v\n", id, ctx.Err())
	
	}

}

func RunTask(){
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()

	for i:=1;i<=3;i++{
		go task(ctx,i)
	}

	time.Sleep(2*time.Second)
	cancel()
	time.Sleep(1*time.Second)
}