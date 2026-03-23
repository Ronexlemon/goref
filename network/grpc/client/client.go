package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/ronexlemon/dsa/network/grpc/shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
	name := flag.String("name", "Yollow", "Name to greet")
	conn,err := grpc.NewClient("localhost:50050",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err !=nil{
		fmt.Println("Error creating a connection",err)
		return
	}
	defer conn.Close()
	c := pb.NewGreaterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloMessage{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())
}