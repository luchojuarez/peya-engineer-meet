package main

import (
	"fmt"
	"time"
)

func timeout() {
	ch1 := make(chan string)
	go callServiceA(ch1)

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Se alcanzo el timeout")
	}
}
