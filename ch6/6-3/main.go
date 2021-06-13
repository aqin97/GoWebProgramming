package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

//帖子的数据结构
type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		{Id: 1, Content: "hello world", Author: "zfq"},
		{Id: 2, Content: "Bonjour Monde", Author: "pierre"},
		{Id: 3, Content: "Hola Mundo", Author: "Pedro"},
		{Id: 4, Content: "hahahahaha", Author: "zfq"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("post.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts)
}
