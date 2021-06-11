//模板的包含动作(涉及到的所有模板都要被解析)
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("t1.html", "t2.html")
	if err != nil {
		fmt.Printf("parse failed")
	}
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
