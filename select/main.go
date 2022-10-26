package main

import (
	"time"
)

func main() {
	timeout()
}

func callServiceA(channel chan string) {
	time.Sleep(3 * time.Second)

	channel <- "api call A ok"
}

func callServiceB(channel chan int) {
	time.Sleep(3 * time.Second)

	channel <- 200
}
