package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "ASCII-WEB/Handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("usage: go run main.go")
		os.Exit(0)
	}
	http.HandleFunc("/", handlers.HomePageHandler)
	fmt.Println("Starting server on http://localhost:5000")
	http.HandleFunc("/ascii-art", handlers.GenerateArt)
	port := ":5000"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Print("Server failed to start succesfully", err)
		return
	}
}
