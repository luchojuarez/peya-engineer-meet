package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")

	result := say("hello")

	fmt.Printf("%d", result)
}

func say(s string) int {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)

		fmt.Println(s)
	}

	return 0
}
