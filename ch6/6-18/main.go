package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
	CreateAt time.Time
}

type Comment struct {
	Id       int
	Content  string
	Author   string `sql:"not null"`
	PostId   int    `sql:"index"`
	CreateAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=admin123 sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "hello world", Author: "zhao fengqin"}
	fmt.Println(post)
	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "good post", Author: "joe"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "zhao fengqin").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
