package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func withErrGroup() {
	g := new(errgroup.Group)
	g.SetLimit(3)

	services := []string{"service-X", "service-Y", "service-Z"}

	for _, service := range services {
		service := service

		g.Go(func() error {
			return apiCallWithErr(service)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(fmt.Sprintf("Error service %s", err.Error()))
	}
}

func apiCallWithErr(service string) error {
	time.Sleep(1 * time.Second)

	if service == "service-Y" {
		return errors.New(fmt.Sprintf("Error in Api Call to %s", service))
	}

	fmt.Printf("Api Call to %s\n\r", service)

	return nil
}
