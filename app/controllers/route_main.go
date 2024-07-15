package controllers

import (
	"net/http"
	"text/template"
)

func Top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/templates/top.html")
	t.Execute(w, "こんにちは")
}
