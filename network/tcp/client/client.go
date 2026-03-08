package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/***
-> The dial function connects to a server
-> The Listen function creates a server
**/

//server

func main(){
	args:=os.Args

	if len(args) == 1{
		fmt.Println("Provide host:port")
		return
	}

	connect := args[1]

	conn,err := net.Dial("tcp",connect)
	if err !=nil{
		fmt.Println("failed to create connection",err)
		return
	}
	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>>")
		text,_:=reader.ReadString('\n')
		fmt.Fprintf(conn,text+"\n")
		message,_:= bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }
	}
}


