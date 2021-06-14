package main

import (
	"encoding/json"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "helloworld",
		Author: Author{
			Id:   2,
			Name: "zhao fengqin",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "good post",
				Author:  "adam",
			},
			{
				Id:      4,
				Content: "have a nice day",
				Author:  "betty",
			},
		},
	}

	jsonFile, err := os.Create("post.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		panic(err)
	}
}
