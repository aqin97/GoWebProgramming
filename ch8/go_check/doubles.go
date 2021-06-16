package main

//用来测试的 测试替身（test doubles）
type FakePost struct {
	Id      int
	Content string
	Author  string
}

func (post *FakePost) fetch(id int) error {
	post.Id = id
	return nil
}

func (post *FakePost) create() error {
	return nil
}

func (post *FakePost) update() error {
	return nil
}

func (post *FakePost) delete() error {
	return nil
}
