package main

import (
	"fmt"
)

func main() {
	// jobs := make(chan int, 5)
	jobs := make(chan int, 5)

	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
				fmt.Println("received more", more)
			} else {
				fmt.Println("received all jobs")
				fmt.Println("over more", more)
				done <- true
				// fmt.Println("done value", done)

				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	// fmt.Println("the done value", done)
	<-done
}
