package main

import (
	"fmt"
	"net/http"
	"text/template"

	"ASCII-WEB/ascii-art1/printingasciipackage"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/ascii-art", GenerateArt)
	fmt.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server failed to start")
		return
	}
}

type ArtDetails struct {
	UserInput  string
	BannerFile string
	Result     string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.ServeFile(w, r, "405.html")
		return
	}

	tmpl.Execute(w, nil)
}

func GenerateArt(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "405.html")
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	text := r.FormValue("userInput")
	banner := r.FormValue("banners")
	fmt.Print(text)

	if text == "" {
		http.ServeFile(w, r, "400.html")
		return
	}
	if banner == "" {
		http.ServeFile(w, r, "400.html")
		return
	}

	result1, err := printingasciipackage.PrintingAscii(text, banner)
	if err != nil {
		http.ServeFile(w, r, "404.html")
		return
	}

	neededData := ArtDetails{
		UserInput:  text,
		BannerFile: banner,
		Result:     result1,
	}

	err = tmpl.Execute(w, neededData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
