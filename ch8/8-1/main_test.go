package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("wrong id, expecting 1 but get ", post.Id)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("skipping encoding for now")
}

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping encoding for now")
	}
	time.Sleep(10 * time.Second)
}
