package main

import (
	"log"
	"net/http"
)

type Gorgonzola struct{}

func NewGorgonzola() *Gorgonzola {
	return &Gorgonzola{}
}

func (g *Gorgonzola) Run() {
	serverConfig := ":8080"

	log.Printf("Starting server at %s", serverConfig)
	if err := http.ListenAndServe(serverConfig, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
