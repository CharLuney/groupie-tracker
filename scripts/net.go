package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Create the web server and waits for incoming user request
func CreateWebsite() {
	fmt.Println("Creating website..", web.Port)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	fmt.Println("Assigning routes..")
	http.HandleFunc("/", Index)
	http.HandleFunc("/filter", Filter)
	http.HandleFunc("/search", Search)
	http.HandleFunc("/inspect", Inspect)
	http.ListenAndServe(web.Port, nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {
	Back()
	ParseTemplate(w)
}

func Filter(w http.ResponseWriter, r *http.Request) {
	filters.MembersMin = r.FormValue("MembersMin")
	filters.MembersMax = r.FormValue("MembersMax")
	filters.CreationDate = r.FormValue("CreationDate")
	filters.FirstAlbum = r.FormValue("FirstAlbum")
	ApplyFilters()
	FilterTemplate(w)
}

func Search(w http.ResponseWriter, r *http.Request) {
	query.Input = r.FormValue("SearchBar")
	fmt.Println("searched for: ", query.Input)
	SearchFor(query.Input)
	SearchTemplate(w)
}

func Inspect(w http.ResponseWriter, r *http.Request) {
	web.Template = "web/inspect.html"
	id, _ := strconv.Atoi(r.FormValue("id"))
	InspectTemplate(w, id)
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

// Parses HTML Template for Filters
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

func SearchTemplate(w http.ResponseWriter) {
	fmt.Println("Parsing template.. ", web.Template)
	tmpl, err := template.ParseFiles(web.Template)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		fmt.Println("Template path: ", web.Template)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artistsSearched)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Parses HTML Template for Inspect
func InspectTemplate(w http.ResponseWriter, id int) {
	fmt.Println("Parsing template.. ", web.Template)
	tmpl, err := template.ParseFiles(web.Template)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		fmt.Println("Template path: ", web.Template)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists[id])
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Back() {
	web.Template = "web/page.html"
}
