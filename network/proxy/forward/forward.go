package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func handleForwardProxy(w http.ResponseWriter,r *http.Request){
	client:= &http.Client{}

	req,err := http.NewRequest(r.Method,r.URL.String(),r.Body)
	if err !=nil{
		fmt.Println("Failed to create new request",err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//copy headers
	for k,v := range r.Header{
		req.Header[k] =v
	}
	//forward req
	resp,err := client.Do(req)
	if err != nil {
        http.Error(w, "Server Error", http.StatusInternalServerError)
        return
    }

	defer resp.Body.Close()
	//copy response header and status
	for k,v := range req.Header{
		w.Header()[k]=v
	}
	w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)

}
func main(){
	http.HandleFunc("/",handleForwardProxy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
