package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
	args:= os.Args

	if len(args) == 1{
		fmt.Println("host:port missing")
		return
	}

	Connect := args[1]
    
	s,err := net.ResolveUDPAddr("udp4",Connect)
	udpConn,err:= net.DialUDP("udp4",nil,s)
	 if err != nil {
                fmt.Println(err)
                return
        }
 fmt.Printf("UDB Server is %s \n",udpConn.RemoteAddr())
 udpConn.Close()

 for{
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">>>")
	test,_:= reader.ReadString('\n')
	data := []byte(test +"\n")
	_,err:=udpConn.Write(data)
	 if strings.TrimSpace(string(data)) == "STOP" {
                        fmt.Println("Exiting UDP client!")
                        return
                }
	if err != nil {
                        fmt.Println(err)
                        return
                }
	buffer := make([]byte, 1024)
                n, _, err := udpConn.ReadFromUDP(buffer)
                if err != nil {
                        fmt.Println(err)
                        return
                }
                fmt.Printf("Reply: %s\n", string(buffer[0:n]))

 }
}