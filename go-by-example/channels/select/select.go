package main

import (
	"fmt"
	"time"
)

func main() {
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if

	// For our example we'll select across two channels.

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "first message"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "second message"
	}()

	for i := 0; i < 2; i++ {
		// Select blocks until a case is ready to run
		select {
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}
}