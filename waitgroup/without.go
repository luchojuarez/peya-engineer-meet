package main

import (
	"fmt"
	"time"
)

func without() {
	services := []string{"service-X", "service-Y", "service-Z"}

	for _, service := range services {
		go apiCall(service)
	}

	fmt.Println("Fin del programa")
}

func apiCall(service string) {
	fmt.Printf("Api Call to %s\n\r", service)
	time.Sleep(1 * time.Second)
}
