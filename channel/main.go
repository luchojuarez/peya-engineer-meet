package main

import (
	"fmt"
	"time"
)

func main() {
	pipe := make(chan string)

	// Anonymous function
	go func(ch chan string) {
		time.Sleep(1 * time.Second)

		ch <- "world"
	}(pipe)

	//go namedFunction(pipe)

	msg := <-pipe

	fmt.Printf("hello %s", msg)
}

func namedFunction(ch chan string) {
	time.Sleep(1 * time.Second)

	ch <- "world"
}
