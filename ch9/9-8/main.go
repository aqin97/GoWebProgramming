package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func thrower(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("threw >>", i)
	}
	wg.Done()
}

func catcher(c chan int) {
	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Println("catched <<", num)
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	wg.Add(2)
	go thrower(c)
	go catcher(c)
	wg.Wait()
}
