package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type ReverseProxy struct{
	backendUrls []*url.URL
	current uint64
}

func NewReverseProxy(backendurls []string)*ReverseProxy{
	urls:= make([]*url.URL,len(backendurls))

	for k,v :=range backendurls{
		parseUrl,err := url.Parse(v)
		if err !=nil{
			log.Fatalf("Invalid URL: %v", err)
		}
		urls[k]= parseUrl
	}
	return &ReverseProxy{backendUrls: urls}
}

func (p ReverseProxy) ServeHTTP(w http.ResponseWriter,r *http.Request){
	index := atomic.AddUint64(&p.current, 1) % uint64(len(p.backendUrls))
    proxy := httputil.NewSingleHostReverseProxy(p.backendUrls[index])
    proxy.ServeHTTP(w, r)

}

func main() {
    backends := []string{"http://localhost:8081", "http://localhost:8082"}
    proxy := NewReverseProxy(backends)
    log.Fatal(http.ListenAndServe(":8080", proxy))
}