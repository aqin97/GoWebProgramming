//Sqlx
package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=admin123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (err error) {
	post := Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post)
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "hello world", AuthorName: "zhao fengqin"}
	err := post.Create()
	fmt.Println(post, err)
}
