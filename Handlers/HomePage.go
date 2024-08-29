package handlers

import (
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		RenderErr(w, http.StatusNotFound, "Page Not Found")

		return
	}
	if r.Method != http.MethodGet {
		RenderErr(w, http.StatusMethodNotAllowed, "Method Not Allowed")

		return
	}
	tmpl.Execute(w, nil)
}
