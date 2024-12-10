package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAPI() {
	// Retrieve API data
	fmt.Println("Getting API")

	artists, err := http.Get(web.Artists)
	if err != nil {
		fmt.Println("Could not get API")
		fmt.Println("Error: ", err.Error())
	}

	locations, err := http.Get(web.Locations)
	if err != nil {
		fmt.Println("Could not get API")
		fmt.Println("Error: ", err.Error())
	}

	dates, err := http.Get(web.Dates)
	if err != nil {
		fmt.Println("Could not get API")
		fmt.Println("Error: ", err.Error())
	}
	relations, err := http.Get(web.Relations)

	if err != nil {
		fmt.Println("Could not get API")
		fmt.Println("Error: ", err.Error())
	}

	// Read API data
	artistsData, err := io.ReadAll(artists.Body)
	if err != nil {
		fmt.Println("Could not read API")
		fmt.Println("Error: ", err.Error())
	}

	locationsData, err := io.ReadAll(locations.Body)
	if err != nil {
		fmt.Println("Could not read API")
		fmt.Println("Error: ", err.Error())
	}

	datesData, err := io.ReadAll(dates.Body)
	if err != nil {
		fmt.Println("Could not read API")
		fmt.Println("Error: ", err.Error())
	}

	relationsData, err := io.ReadAll(relations.Body)
	if err != nil {
		fmt.Println("Could not read API")
		fmt.Println("Error: ", err.Error())
	}

	// Convert API data
	json.Unmarshal(artistsData, &artists)
	json.Unmarshal(locationsData, &locations)
	json.Unmarshal(datesData, &dates)
	json.Unmarshal(relationsData, &relations)
}
