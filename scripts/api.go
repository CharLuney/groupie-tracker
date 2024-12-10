package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAll() {
	GetArtists()
	GetLocations()
	GetDates()
	GetRelations()
	fmt.Println("DEBUG: ", artists[0].Name)
}

func GetAPI(API string) []byte {
	// Retrieve API data
	groupie, err := http.Get(web.Artists)

	if err != nil {
		fmt.Println("Could not get API")
		fmt.Println("Error: ", err.Error())
	}

	// Read API data
	groupieData, err := io.ReadAll(groupie.Body)

	if err != nil {
		fmt.Println("Could not read API")
		fmt.Println("Error: ", err.Error())
	}

	return groupieData

}

func GetArtists() {
	fmt.Println("Getting API.. Artists")
	data := GetAPI(web.Artists)
	json.Unmarshal(data, &artists)
}

func GetLocations() {
	fmt.Println("Getting API.. Locations")
	data := GetAPI(web.Locations)
	json.Unmarshal(data, &locations)
}

func GetDates() {
	fmt.Println("Getting API.. Dates")
	data := GetAPI(web.Dates)
	json.Unmarshal(data, &dates)
}

func GetRelations() {
	fmt.Println("Getting API.. Relations")
	data := GetAPI(web.Relations)
	json.Unmarshal(data, &relations)
}
