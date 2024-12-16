package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// Create the web server and waits for incoming user request
func CreateWebsite() {
	fmt.Println("Creating website..", web.Port)
	http.HandleFunc("/", Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.ListenAndServe(web.Port, nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {
	ParseTemplate(w)
}

// Parses HTML Template
func ParseTemplate(w http.ResponseWriter) {
	fmt.Println("Parsing template.. ", web.Template)
	tmpl, err := template.ParseFiles(web.Template)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		fmt.Println("Template path: ", web.Template)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
