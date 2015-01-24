package gorgonzola

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Gorgonzola struct {
	c       *Config
	storage Storage
}

func NewGorgonzola() *Gorgonzola {
	return &Gorgonzola{
		c:       NewConfig(),
		storage: NewDatastore(),
	}
}

type httpHandler func(http.ResponseWriter, *http.Request) error

func (g *Gorgonzola) setHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/", httpHandler(g.indexHandler).ServeHTTP).Methods("GET")
	r.HandleFunc("/add.html", httpHandler(g.addHandler).ServeHTTP).Methods("GET", "POST")
	http.Handle("/", r)
}

func (g *Gorgonzola) Run() {
	g.setHandlers()
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
