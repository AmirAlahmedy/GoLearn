package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// A nonblocking recevie.
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	// A nonblocking send works similarly.
	// Here msg cannot be sent to the channel, because the channel has no buffer and there is no receiver.
	msg := "hi"
	select {
		case messages <- msg:
			fmt.Println("sent message:", msg)
		default: 
		    fmt.Println("failed to send message:", msg)
	}

	// A multi-way non-blocking select.
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}