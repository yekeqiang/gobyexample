package main

import (
	"fmt"
	// "time"
)

func main() {
	// messages := make(chan string) //如果不加缓存的话，就全部会选择defalut
	messages := make(chan string, 1) //加了缓存的话，会选择对应的
	signals := make(chan bool)

	// go func() {
	// time.Sleep(time.Second * 3)
	messages <- "test"
	// }()
	// messages <- "test"
	// time.Sleep(time.Second * 1)
	select {
	// case <-time.After(time.Second * 2):
	// fmt.Println("received messages")
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// go func() {
	msg := "hi world"
	// }()
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
