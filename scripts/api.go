package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Get all APIs
func GetAll() {
	GetArtists()
	GetDates()
	GetLocations()
	GetRelations()
	SetIDs()
}

// Gets the API from the provided URL and returns its data
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

// Get artists branch of API
func GetArtists() {
	fmt.Println("Getting API.. Artists")
	data := GetAPI(web.Artists)
	json.Unmarshal(data, &artists)
}

// Get locations branch of API
func GetLocations() {
	fmt.Println("Getting API.. Locations")
	data := GetAPI(web.Locations)
	json.Unmarshal(data, &artists)
}

// Get dates branch of API
func GetDates() {
	fmt.Println("Getting API.. Dates")
	data := GetAPI(web.Dates)
	json.Unmarshal(data, &artists)
}

// Get relations branch of API
func GetRelations() {
	fmt.Println("Getting API.. Relations")
	data := GetAPI(web.Relations)
	json.Unmarshal(data, &artists)
}

// Sets IDs to each artist
func SetIDs() {
	fmt.Println("Assigning all IDs..")
	for i := 0; i < len(artists); i++ {
		artists[i].ID = strconv.Itoa(i)
	}
}
