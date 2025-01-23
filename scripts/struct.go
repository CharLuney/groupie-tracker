package api

var web Website

var filters Filters
var regex Regex
var query Query

var artists []Artist
var artistsFilterted []Artist
var artistsSearched []Artist

type Website struct {
	Template string // Path to template file to parse
	Port     string // Port used for the server

	Artists   string // API Url to Artists branch
	Locations string // API Url to Locations branch
	Dates     string // API Url to Dates branch
	Relations string // API Url to Relations branch
	Map       string // API Url to OpenStreetMap (OSM)
}

type Filters struct {
	CreationDate    string // Value of creation date filter
	FirstAlbum      string // Value of first album filter
	MembersMin      string // Minimum amount of members filter
	MembersMax      string // Maximum amount of members filter
	ConcertLocation string // Value of concert location filter
	Verified        bool   // True if the artist passed all the filters
}

type Regex struct {
	Name string // Type string
	Date string // Type date
	Year string // Type year
}

type Query struct {
	Input     string // Search bar's user input
	IsName    bool   // True if query type is a string
	IsDate    bool   // True if query type is a date
	IsYear    bool   // True if query type is a year
	IsInvalid bool   // True if query type is a invalid
	LookedFor bool   // Determinates if the current artist is looked for
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
