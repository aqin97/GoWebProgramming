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

}

func (post *Post) delete() error {

}
