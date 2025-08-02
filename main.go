package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	ART    string
	ERR    string
	Banner string
}

var latestArt PageData

// max lenght
// 404 page not found
// unsupported character

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Invalid arguments")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			if r.Method != http.MethodGet {
				http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
				return
			}
			if err := tmpl.Execute(w, latestArt); err != nil {
				http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
			}

		case "/ascii-art":
			if r.Method != http.MethodPost {
				http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
				return
			}

			banner := r.FormValue("banner")
			input := r.FormValue("inputText")

			if banner == "" || input == "" || len(input) > 900 {
				http.Error(w, "Bad request 400", http.StatusBadRequest)
				return
			}

			art, err := run(input, banner)
			if err != nil {
				http.Error(w, "Bad request 400", http.StatusBadRequest)
				return
			}

			latestArt = PageData{ART: art, Banner: banner}

			if err := tmpl.Execute(w, latestArt); err != nil {
				http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
			}

		default:
			http.NotFound(w, r)
		}
	})

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
