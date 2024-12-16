package api

var web Website

var data []Data

type Website struct {
	Template string
	Port     string

	Artists   string
	Locations string
	Dates     string
	Relations string
}

type Data struct {
	ID           string   `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
	Locations    []string `json:"locations"`
	ConcertDates []string `json:"dates"`
	Relations    []string `json:"relations"`
}
