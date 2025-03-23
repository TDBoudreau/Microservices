package main

import (
	"fmt"
	"frontend/cmd/web/templates"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	// partials := []string{
	// 	"./cmd/web/templates/base.layout.gohtml",
	// 	"./cmd/web/templates/header.partial.gohtml",
	// 	"./cmd/web/templates/footer.partial.gohtml",
	// }

	partials := []string{
		"base.layout.gohtml",
		"header.partial.gohtml",
		"footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, t)
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFS(templates.TemplatesFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		BrokerURL string
	}

	data.BrokerURL = os.Getenv("BROKER_URL")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
