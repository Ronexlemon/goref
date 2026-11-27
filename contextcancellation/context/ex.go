package context

import (
	"context"
	"fmt"
)

//tracing through layers

func queryDB(ctx context.Context){
	traceID:=ctx.Value("traceID").(string)
	fmt.Println("DB trace:", traceID)
}

func callRPC(ctx context.Context){
	ctx=context.WithValue(ctx,"traceID","trace-1234")
	queryDB(ctx)
}

func RunTracingLayer(){
	ctx:=context.Background()
	callRPC(ctx)
}