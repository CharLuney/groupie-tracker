package api

var web Website
var filters Filters

var artists []Artist
var artistsFilterted []Artist

type Website struct {
	Template string
	Port     string

	Artists   string
	Locations string
	Dates     string
	Relations string
}

type Filters struct {
	CreationDate    string
	FirstAlbum      string
	MembersAmount   string
	ConcertLocation string
	Satisfied       bool
}

type Artist struct {
	ID           string
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
	Locations    []string `json:"locations"`
	ConcertDates []string `json:"dates"`
	Relations    []string `json:"relations"`
}
