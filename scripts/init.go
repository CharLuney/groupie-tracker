package api

import "fmt"

func Init() {
	fmt.Println("Initializing..")

	web = Website{
		Template:  "web/page.html",
		Port:      ":8080",
		Artists:   "https://groupietrackers.herokuapp.com/api/artists",
		Locations: "https://groupietrackers.herokuapp.com/api/locations",
		Dates:     "https://groupietrackers.herokuapp.com/api/dates",
		Relations: "https://groupietrackers.herokuapp.com/api/relations",
	}

	regex = Regex{
		Name: "[a-z]+$",
		Date: "[0-9]{2}/[0-9]{2}/[0-9]{2}$",
		Year: "[0-9]{4}$",
	}

	GetAll()
	fmt.Println(artists[0])

	CreateWebsite()
}
