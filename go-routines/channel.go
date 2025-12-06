package main

import (
	"fmt"
	"time"
)

func main() {
	channelWithSelectClause()
}

func channelWithSelectClause() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "func one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "func two"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}

	}

}

func basicChannelExample() {
	message := make(chan string)

	go func(msg string) {
		message <- msg
	}("ping")

	receiver := <-message

	fmt.Println(receiver)
}
