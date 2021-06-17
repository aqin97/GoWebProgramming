package main

import (
	"fmt"
	"time"
)

func printNumbers2(w chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}

	w <- 1
}

func printLetters2(w chan int) {
	for i := 'a'; i < 'a'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}

	w <- 1
}

func main() {
	w1 := make(chan int)
	w2 := make(chan int)
	go printLetters2(w1)
	go printNumbers2(w2)
	<-w1
	<-w2
}
