package api

var web Website

var artists []Artist
var locations []Location
var dates []Date
var relations []Relation

type Website struct {
	Template  string
	Artists   string
	Locations string
	Dates     string
	Relations string
}

type Artist struct {
	ID           string   `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
}

type Location struct {
	Locations []string
}

type Date struct {
	ConcertDates []string
}

type Relation struct {
	Relations []string
}
