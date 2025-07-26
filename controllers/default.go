package controllers

import (
	"html/template"
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/index.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
