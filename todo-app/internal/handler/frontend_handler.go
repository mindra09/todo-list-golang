package handler

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Heading string
	Name    string
	Items   []string
}
type FrontEndHandler struct {
}

func NewFrontEndHandler() *FrontEndHandler {
	return &FrontEndHandler{}
}

func (h *FrontEndHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./frontend/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
