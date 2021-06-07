package main

import (
	"fmt"
	"net/http"
)

func main() {
	//处理静态文件
	mux := http.NewServeMux()
	//创建一个能为指定目录里的静态文件服务的处理器，并将这个处理器作为参数传递给多路复用器的Handle函数
	/*files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	*/
	//发送至根URL的请求重定向到index这个处理器函数
	mux.HandleFunc("/", index)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
