package context

import (
	"context"
	"fmt"
	"log"
	"time"
)

type key string
var traceID key ="traceID"
type OrderResponse struct {
    TraceID     string `json:"trace_id"`
    OrderID     string `json:"order_id"`
    Status      string `json:"status"`
    StockStatus string `json:"stock_status"`
    Error       string `json:"error,omitempty"`
}
type result struct {
        status string
        stock  string
        err    error
    }

func queryDBForStatus(ctx context.Context,orderID string)(string ,error){
	select{
	case <-time.After(600 *time.Millisecond):
		return "shipped",nil
	case <-ctx.Done():
		return "",ctx.Err()
	}
}

func CallStockInventory(ctx context.Context,orderID string)(string ,error){
	select{
	case <-time.After(400*time.Millisecond):
		return "in_stock", nil
	case <-ctx.Done():
		return "",ctx.Err()

	}
}

func RunStockInventory(orderId string ){

	ctx,cancel:=context.WithTimeout(context.Background(),1500*time.Millisecond)
	defer cancel()
 traceId := fmt.Sprintf("req-%d", time.Now().UnixNano())
	ctx=context.WithValue(ctx,traceID,traceId)

	select{
	case <-ctx.Done():
		resp:=OrderResponse{TraceID: traceId,OrderID: orderId,Error: "client ditched"}
		//return n response
		log.Printf("Canceled [traceID=%s]: %v", resp.TraceID, ctx.Err())
        return
	default:

	}

	ch:=make(chan result,1 )

	go func() {
		status,err:=queryDBForStatus(ctx,orderId)
		if err !=nil{
			ch<-result{err: err}
			return 
		}
		stock,err:=CallStockInventory(ctx,orderId)
		if err !=nil{
			ch <-result{err: err}
			return
		}
		ch <-result{status: status,stock: stock,err: nil}
	}()

	var resp OrderResponse


	select{
	case res:=<-ch:
		if res.err !=nil{
			resp = OrderResponse{TraceID: traceId, OrderID: orderId, Error: res.err.Error()}
			 log.Printf("Failed [traceID=%s]: %v and result %v", traceID, res.err,resp)
            return
		}
		resp = OrderResponse{TraceID: traceId, OrderID: orderId, Status: res.status, StockStatus: res.stock}
    case <-ctx.Done():
		resp = OrderResponse{TraceID: traceId, OrderID: orderId, Error: "Timed out"}
		log.Printf("Timeout [traceID=%s]: %v and result %v", traceID, ctx.Err(),resp)
        return
	}

	log.Printf("Done [traceID=%s] and response %v", traceID,resp)

}