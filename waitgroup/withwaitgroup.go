package main

import (
	"fmt"
	"sync"
	"time"
)

func withWaitGroup() {
	var wg sync.WaitGroup

	services := []string{"service-X", "service-Y", "service-Z"}

	for _, service := range services {
		wg.Add(1)

		go apiCallWithGroup(&wg, service)
	}

	wg.Wait() // Espera hasta que se ejecuten las goroutines

	fmt.Println("Fin del programa")
}

func apiCallWithGroup(wg *sync.WaitGroup, service string) {
	defer wg.Done() // Disminuye en uno la cantidad de goroutines a esperar

	fmt.Printf("Api Call to %s\n\r", service)
	time.Sleep(1 * time.Second)
}
