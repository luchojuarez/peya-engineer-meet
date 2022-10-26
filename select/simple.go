package main

import "fmt"

func simple() {
	ch1 := make(chan string)
	go callServiceA(ch1)

	ch2 := make(chan int)
	go callServiceB(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(fmt.Sprintf("service B: %d", val2))
	}
}
