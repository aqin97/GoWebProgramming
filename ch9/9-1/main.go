package main

import (
	"fmt"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'a'; i < 'a'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printLetters1()
	printNumbers1()
}

func goprint1() {
	go printLetters1()
	go printNumbers1()
}

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters2() {
	for i := 'a'; i < 'a'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
}

func goprint2() {
	go printLetters2()
	go printNumbers2()
}

func main() {
	print1()
	goprint1()
	goprint2()
}
