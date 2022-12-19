package main

import (
	"log"

	"github.com/Vadia1337/Http-restAPI-Golang/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
