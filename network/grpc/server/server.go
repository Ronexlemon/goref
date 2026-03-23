package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ronexlemon/dsa/network/grpc/shared/proto"
	"google.golang.org/grpc"
)

type ServerGreate struct{
	pb.UnimplementedGreaterServer
}

func (s *ServerGreate) SayHello(_ context.Context,input *pb.HelloMessage)(*pb.HelloReply,error){
	return &pb.HelloReply{Name: input.GetName()},nil
}

func main(){
	lis,err := net.Listen("tcp",":50050")
	if err !=nil{
		fmt.Println("failed to listen",err)
		return
	}
	s:= grpc.NewServer()
	pb.RegisterGreaterServer(s,&ServerGreate{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}