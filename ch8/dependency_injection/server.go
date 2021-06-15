package main

import (
	"database/sql"
	"net/http"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, t Text) error {

}

func handlePost(w http.ResponseWriter, r *http.Request, t Text) error {

}

func handlePut(w http.ResponseWriter, r *http.Request, t Text) error {

}

func handleDelete(w http.ResponseWriter, r *http.Request, t Text) error {

}

func main() {
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=admin123 sslmode=disable")
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	post := Post{Db: db}
	http.HandleFunc("/post/", handleRequest(&post))

	server.ListenAndServe()
}
