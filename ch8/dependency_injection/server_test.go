package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	requset, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, requset)

	if writer.Code != http.StatusOK {
		t.Errorf("response code id %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("cannot retrieve post")
	}
}

func TestPutPost(t *testing.T) {
	mux := http.NewServeMux()
	post := &FakePost{}
	mux.HandleFunc("/post/", handleRequest(post))

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"updated post", "author":"zhao fengqin"}`)
	requset, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, requset)

	if writer.Code != http.StatusOK {
		t.Errorf("response code id %v", writer.Code)
	}
	if post.Content != "updated post" {
		t.Error("Content is not correct", post.Content)
	}
}
