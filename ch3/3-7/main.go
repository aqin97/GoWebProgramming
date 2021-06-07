package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

//处理器和处理器实现的ServeHTTP方法
type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func hellozfq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello zfq")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	http.HandleFunc("/zfq", hellozfq)

	http.Handle("/aaa", http.HandlerFunc(hellozfq))

	server.ListenAndServe()
}
