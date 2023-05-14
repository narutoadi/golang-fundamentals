// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine

package main

import "fmt"

func main() {
	messages := make(chan string)

	// channel <-
	// sends a value into channel
	go func() { messages <- "ping " }()

	// messages <- "ping "

	// <-channel
	// receives a value from the channel
	msg := <-messages
	fmt.Println(msg)
}
