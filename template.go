package main

import (
	"html/template"
	"net/http"
)

type templateData map[string]interface{}

type Template struct {
	data   templateData
	layout *template.Template
	w      http.ResponseWriter
}

func NewTemplate(w http.ResponseWriter) *Template {
	return &Template{
		data:   make(templateData),
		layout: template.New("layout.html"),
		w:      w,
	}
}

func (t Template) render(filenames ...string) error {
	return template.Must(t.layout.ParseFiles(filenames...)).Execute(t.w, t.data)
}
