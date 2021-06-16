package main

import (
	"fmt"
)

func callerA(c chan<- string) {
	c <- "hello world"
}

func callerB(c chan<- string) {
	c <- "nihao"
}

func main() {
	a := make(chan string)
	b := make(chan string)
	go callerA(a)
	go callerB(b)
	for i := 0; i < 5; i++ {
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("default")
		}
	}
}
