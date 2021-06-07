//串联处理器，本质上和串联处理器函数区别很小
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h.ServeHTTP).Pointer()).Name()
		fmt.Println("Handler called - " + name)
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	helloHandler := HelloHandler{}
	http.Handle("/hello/", log(&helloHandler))
	server.ListenAndServe()
}
