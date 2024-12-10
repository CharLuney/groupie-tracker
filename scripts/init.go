package api

import "fmt"

func Init() {
	fmt.Println("Initializing")

	web = Website{
		Template:  "static/page.html",
		Artists:   "https://groupietrackers.herokuapp.com/api/artists",
		Locations: "https://groupietrackers.herokuapp.com/api/locations",
		Dates:     "https://groupietrackers.herokuapp.com/api/dates",
		Relations: "https://groupietrackers.herokuapp.com/api/relations",
	}

	CreateWebsite()

}
