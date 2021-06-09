//通过WriteHeader方法将状态码写到响应中
package main

import (
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

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeheaderExample)

	server.ListenAndServe()
}
