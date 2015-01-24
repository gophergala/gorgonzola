package gorgonzola

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (g *Gorgonzola) indexHandler(w http.ResponseWriter, r *http.Request) error {
	tm := NewTemplate(w)
	jobs, err := g.storage.GetJobs(r)
	if err != nil {
		return err
	}
	tm.set("jobs", jobs)
	return tm.render("templates/layout.html", "templates/index.html")
}

func (g *Gorgonzola) jobHandler(w http.ResponseWriter, r *http.Request) error {
	tm := NewTemplate(w)
	vars := mux.Vars(r)
	job, err := g.storage.GetJob(r, vars["key"])
	if err != nil {
		return err
	}
	tm.set("job", job)
	return tm.render("templates/layout.html", "templates/job.html")
}

func (g *Gorgonzola) addHandler(w http.ResponseWriter, r *http.Request) error {
	tm := NewTemplate(w)
	if r.Method == "POST" && r.FormValue("url") != "" {
		id, err := g.storage.AddURL(r, r.FormValue("url"))
		if err != nil {
			return err
		}
		http.Redirect(w, r, "/task/"+id, http.StatusFound)
	}
	return tm.render("templates/layout.html", "templates/add.html")
}
