package main

import (
	"rbac-demo/cmd/app/options"
)

func main() {

	s := options.NewOptions()

	if err := s.Complete(); err != nil {
		panic(err)
	}

	s.Run()
}
