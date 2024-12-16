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
}

func GetAPI(API string) []byte {
	// Retrieve API data
	groupie, err := http.Get(API)

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
	json.Unmarshal(data, &data)
}

func GetLocations() {
	fmt.Println("Getting API.. Locations")
	data := GetAPI(web.Locations)
	json.Unmarshal(data, &data)
}

func GetDates() {
	fmt.Println("Getting API.. Dates")
	data := GetAPI(web.Dates)
	json.Unmarshal(data, &data)
}

func GetRelations() {
	fmt.Println("Getting API.. Relations")
	data := GetAPI(web.Relations)
	json.Unmarshal(data, &data)
}
