package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// Create the web server and waits for incoming user request
func CreateWebsite() {
	fmt.Println("Creating website..", web.Port)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", Index)
	http.HandleFunc("/filter", Filter)
	http.HandleFunc("/inspect", Inspect)
	http.ListenAndServe(web.Port, nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {
	ParseTemplate(w)
}

func Filter(w http.ResponseWriter, r *http.Request) {
	filter.MembersAmount = r.FormValue("Members")
	fmt.Println(filter.MembersAmount)
	FilterMembers()
	FilterTemplate(w)
}

func Inspect(w http.ResponseWriter, r *http.Request) {
	web.Template = "web/inspect.html"
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

	err = tmpl.Execute(w, artists)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Parses HTML Template
func FilterTemplate(w http.ResponseWriter) {
	fmt.Println("Parsing template.. ", web.Template)
	tmpl, err := template.ParseFiles(web.Template)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		fmt.Println("Template path: ", web.Template)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artistsFilterted)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
