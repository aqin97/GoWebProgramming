package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	_ "github.com/lib/pq"
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
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	err = t.fetch(id)
	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(t, "", "	")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return nil
}

func handlePost(w http.ResponseWriter, r *http.Request, t Text) error {
	len := r.ContentLength
	body := make([]byte, len)
	_, err := r.Body.Read(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, t)
	if err != nil {
		return err
	}
	err = t.create()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func handlePut(w http.ResponseWriter, r *http.Request, t Text) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	err = t.fetch(id)
	if err != nil {
		return err
	}
	len := r.ContentLength
	body := make([]byte, len)
	_, err = r.Body.Read(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, t)
	if err != nil {
		return err
	}
	err = t.update()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func handleDelete(w http.ResponseWriter, r *http.Request, t Text) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	err = t.fetch(id)
	if err != nil {
		return err
	}
	err = t.delete()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
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
