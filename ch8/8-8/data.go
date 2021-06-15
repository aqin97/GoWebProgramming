package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

//通过初始化函数连接到数据库
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=admin123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (Post, error) {
	post := Post{}
	err := Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return post, err
}

func (post *Post) Create() error {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return err
}

func (post *Post) Update() error {
	_, err := Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return err
}

func (post *Post) Delete() error {
	_, err := Db.Exec("delete from posts where id = $1", post.Id)
	return err
}
