package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Gorgonzola struct {
	c *Config
}

func NewGorgonzola() *Gorgonzola {
	return &Gorgonzola{
		c: NewConfig(),
	}
}

type httpHandler func(http.ResponseWriter, *http.Request) error

func (g *Gorgonzola) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", httpHandler(g.indexHandler).ServeHTTP).Methods("GET")

	http.Handle("/", r)

	log.Printf("Starting server at %s", g.c.Server)
	if err := http.ListenAndServe(g.c.Server, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}

func (fn httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		httpError, ok := err.(HTTPError)
		if ok {
			http.Error(w, httpError.Message, httpError.Code)
			return
		}
		// Default to 500 Internal Server Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
