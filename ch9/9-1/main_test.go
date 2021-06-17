package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T) {
	print1()
}

func TestGoPrint1(t *testing.T) {
	goprint1()
	time.Sleep(1 * time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
	goprint2()
	time.Sleep(1 * time.Millisecond)
}
