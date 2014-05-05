package main

import (
	"fmt"
)

var a string
var c = make(chan string)

func f() {
	a = "hello, world"
	// <-c
	c <- a
	// c <- 0
}

func main() {
	go f()
	// c <- 0
	// <-c
	fmt.Print(<-c)
}
