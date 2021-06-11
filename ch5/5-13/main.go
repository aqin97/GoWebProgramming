package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("t.html")
	if err != nil {
		fmt.Printf("parse failed")
	}
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
