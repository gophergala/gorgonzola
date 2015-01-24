package gorgonzola

import (
	"net/http"
)

func (g *Gorgonzola) indexHandler(w http.ResponseWriter, r *http.Request) error {
	tm := NewTemplate(w)
	jobs, err := g.storage.GetJobs()
	if err != nil {
		return err
	}
	tm.set("jobs", jobs)
	return tm.render("templates/layout.html", "templates/index.html")
}

func (g *Gorgonzola) addHandler(w http.ResponseWriter, r *http.Request) error {
	tm := NewTemplate(w)
	if r.Method == "POST" && r.FormValue("url") != "" {
		id, err := g.storage.AddURL(r.FormValue("url"))
		if err != nil {
			return err
		}
		http.Redirect(w, r, "/task/"+id, http.StatusFound)
	}
	return tm.render("templates/layout.html", "templates/add.html")
}
