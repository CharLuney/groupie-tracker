package api

var web Website
var filters Filters

var artists []Artist
var artistsFilterted []Artist

var locations []Locations
var dates []Dates
var relations []Relations

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
	MembersMin      string
	MembersMax      string
	ConcertLocation string
	Satisfied       bool
}

type Artist struct {
	ID           string
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationdate"`
	FirstAlbum   string              `json:"firstalbum"`
	Locations    []string            `json:"locations"`
	ConcertDates []string            `json:"dates"`
	Relations    map[string][]string `json:"relations"`
}

type Locations struct {
	Location []string `json:"locations"`
}

type Dates struct {
	Date []string `json:"dates"`
}

type Relations struct {
	Relation map[string][]string `json:"relations"`
}
