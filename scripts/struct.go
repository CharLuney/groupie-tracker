package api

var web Website
var data []Data

type Website struct {
	Template string
	YnovAPI  string
}

type Data struct {
	Artist      Artist
	Location    Location
	ConcertDate ConcertDate
	Relation    Relation
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

type ConcertDate struct {
	ConcertDates []string
}

type Relation struct {
	Relations []string
}
