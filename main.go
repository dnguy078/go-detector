package main

import (
	"log"

	"github.com/dnguy078/go-detector/daemon"
)

func main() {
	s, err := daemon.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Start(); err != nil {
		log.Fatalf("unable to start server: %s", err)
	}
}
