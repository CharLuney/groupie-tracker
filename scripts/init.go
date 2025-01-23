package api

import (
	"fmt"
)

func Init() {
	fmt.Println("Initializing..")

	// Initializes server properties and API URLs
	web = Website{
		Template:  "web/page.html",
		Port:      ":8080",
		Artists:   "https://groupietrackers.herokuapp.com/api/artists",
		Locations: "https://groupietrackers.herokuapp.com/api/locations",
		Dates:     "https://groupietrackers.herokuapp.com/api/dates",
		Relations: "https://groupietrackers.herokuapp.com/api/relations",
		Map:       "https://api.openstreetmap.org/api/0.6/map?bbox=-0.1275,51.5072,0.1275,51.5078",
	}

	// Initializes regexs for each possible input
	regex = Regex{
		Name: "[a-zA-Z]$",
		Date: "[0-9]{2}-[0-9]{2}-[0-9]{4}$",
		Year: "[0-9]{4}$",
	}

	GetAll()
	CreateWebsite()
}
