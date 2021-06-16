package main

func (post *Post) fetch(id int) error {
	err := post.Db.QueryRow("select id, content, author from where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return err
}

func (post *Post) create() error {
	err := post.Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return err
}

func (post *Post) update() error {
	_, err := post.Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Content, post.Author, post.Id)
	return err
}

func (post *Post) delete() error {
	_, err := post.Db.Exec("delete from post where id = $1", post.Id)
	return err
}
