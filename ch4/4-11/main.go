//编写json输出
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>hello world</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeheaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "no such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "zhao feng qin",
		Threads: []string{"first", "second", "third"},
	}
	jsonBytes, _ := json.Marshal(post)
	w.Write(jsonBytes)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeheaderExample)
	http.HandleFunc("/header", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
