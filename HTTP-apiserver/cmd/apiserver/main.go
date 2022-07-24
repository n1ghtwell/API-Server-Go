package main

import (
	"log"
)

var (
	configPath string
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
