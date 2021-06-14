package main

//创建xml文件

import (
	"encoding/xml"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
	//Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "hello world",
		Author: Author{
			Id:   "2",
			Name: "zhao fengqin",
		},
	}

	output, err := xml.MarshalIndent(&post, "", "	")
	if err != nil {
		panic(err)
	}
	//err = ioutil.WriteFile("post.xml", output, 0644)
	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		panic(err)
	}
}
