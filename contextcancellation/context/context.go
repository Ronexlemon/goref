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