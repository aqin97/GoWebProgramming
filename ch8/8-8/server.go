package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/post/", handleRequest)

	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handPost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(&post, "", "	")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return err
}

func handPost(w http.ResponseWriter, r *http.Request) error {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var post Post
	err := json.Unmarshal(body, &post)
	if err != nil {
		return err
	}
	err = post.Create()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func handlePut(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}

	post, err := retrieve(id)
	if err != nil {
		return err
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	err = json.Unmarshal(body, &post)
	if err != nil {
		return err
	}
	err = post.Update()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func handleDelete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	err = post.Delete()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
