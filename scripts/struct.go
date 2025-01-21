package api

var web Website
var filters Filters
var regex Regex
var query Query

var artists []Artist
var artistsFilterted []Artist
var artistsSearched []Artist

// var locations []IndexLocations

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
	Verified        bool
}

type Regex struct {
	Name string
	Date string
	Year string
}

type Query struct {
	Input     string
	IsName    bool
	IsDate    bool
	IsYear    bool
	IsInvalid bool
	LookedFor bool
}

type Artist struct {
	ID           string
	Image        string           `json:"image"`
	Name         string           `json:"name"`
	Members      []string         `json:"members"`
	CreationDate int              `json:"creationdate"`
	FirstAlbum   string           `json:"firstalbum"`
	Locations    []IndexLocations `json:"locations"`
	ConcertDates []IndexDates     `json:"dates"`
	Relations    []IndexRelations `json:"relations"`
}

type IndexLocations struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

type IndexDates struct {
	Date []Dates `json:"index"`
}

type Dates struct {
	Date []string `json:"dates"`
}

type IndexRelations struct {
	Date []string `json:"index"`
}

type Relations struct {
	Relation map[string][]string `json:"relations"`
}
