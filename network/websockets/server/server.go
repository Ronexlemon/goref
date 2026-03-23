package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request)bool{return true},
}

func handleConnection(w http.ResponseWriter,r *http.Request){
	conn,err := upgrader.Upgrade(w,r,nil)
	if err !=nil{
		fmt.Println("Error creating a connection ",err)
		return
	}
	defer conn.Close()

	for{
	typeData,message,err := conn.ReadMessage()
	if err !=nil{
		fmt.Println("Failed to read message",err)
		break
	}
	fmt.Println("The typeData",typeData)
	 fmt.Printf("Received: %s\\n", message)

	if err:= conn.WriteMessage(websocket.TextMessage,message); err !=nil{
		fmt.Println("Error writing message:", err)
          break
	}

	}
}

func main(){
	http.HandleFunc("/ws",handleConnection)
	fmt.Println("WebSocket server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
       fmt.Println("Error starting server:", err)
    }

}
