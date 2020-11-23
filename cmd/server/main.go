package main

import (
	"log"

	"github.com/MishaNiki/prob-gateway/internal/server"
)

func main() {
	s := server.NewMainServer()

	if err := s.Configure(); err != nil {
		log.Fatal(err)
	}

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
