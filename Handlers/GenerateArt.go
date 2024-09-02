package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"ASCII-WEB/ascii-art1/printingasciipackage"
)

type ArtDetails struct {
	UserInput  string
	BannerFile string
	Result     string
}

func GenerateArt(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		RenderErr(w, http.StatusInternalServerError, "Failed to load result template")
		return
	}
	if r.Method != http.MethodPost {
		RenderErr(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	if r.URL.Path != "/ascii-art" {
		RenderErr(w, http.StatusNotFound, "Page Not Found")
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	text := r.FormValue("userInput")
	banner := r.FormValue("banners")

	if text == "" || banner == "" {
		RenderErr(w, http.StatusBadRequest, "Banner selection is missing")
		return
	}

	result1, err := printingasciipackage.PrintingAscii(text, banner)
	fmt.Println(err)
	if err != nil {
		if err.Error() == "Character not supported" {
			RenderErr(w, http.StatusBadRequest, "Character not supported")
			return
		}
		RenderErr(w, http.StatusInternalServerError, "Oops! something happened contact admin")
		return
	}

	neededData := ArtDetails{
		UserInput:  text,
		BannerFile: banner,
		Result:     result1,
	}

	err = tmpl.Execute(w, neededData)
	if err != nil {
		RenderErr(w, http.StatusInternalServerError, "Failed to render result page")
		return
	}
}
