package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

//创建测试套件
type PostTestSuite struct{}

//通过初始化函数注册测试套件
func init() {
	Suite(&PostTestSuite{})
}

func TestT(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()
	requset, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, requset)

	c.Check(writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}
