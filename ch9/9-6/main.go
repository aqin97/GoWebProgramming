package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2() {
	for i := 'a'; i < 'a'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func goprint2() {
	go printLetters2()
	go printNumbers2()
}

func main() {
	wg.Add(2)
	goprint2()
	wg.Wait()
}
