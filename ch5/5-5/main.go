package main

//模板的迭代动作， 传过去的是 数组，切片，映射，通道等类型 类型

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		fmt.Println("parse failed")
	}
	data := []string{"Mons", "Tues", "Wed", "Thu", "Fri", "Sat", "Sun"}

	t.Execute(w, data)
}

func process2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		fmt.Println("parse failed")
	}
	var data []int
	t.Execute(w, data)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)

	server.ListenAndServe()
}
