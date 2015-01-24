package main

import (
	"net/http"
)

func (g *Gorgonzola) indexHandler(w http.ResponseWriter, r *http.Request) error {
	tm := NewTemplate(w)
	return tm.render("templates/layout.html", "templates/index.html")
}
