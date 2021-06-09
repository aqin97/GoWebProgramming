package main

//模板的条件动作， 传过去的是 bool 类型

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		fmt.Println("parse failed")
	}

	rand.Seed(time.Now().UnixNano())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
