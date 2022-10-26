package main

import (
	"fmt"
	"time"
)

func withDefault() {
	ch1 := make(chan string)
	go callServiceA(ch1)

	ch2 := make(chan int)
	go callServiceB(ch2)

	for {
		time.Sleep(500 * time.Millisecond)

		select {
		case val1 := <-ch1:
			fmt.Println(val1)
		case val2 := <-ch2:
			fmt.Println(fmt.Sprintf("service B: %d", val2))
		default:
			fmt.Println("channels no listos")
		}
	}
}
