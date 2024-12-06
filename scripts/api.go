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
	groupie, err := http.Get("https://groupietrackers.herokuapp.com/api")

	if err != nil {
		fmt.Println("Could not get API")
		fmt.Println("Error: ", err.Error())
	}

	// Read API data
	groupieData, err := io.ReadAll(groupie.Body)

	// Debug
	fmt.Println(string(groupieData))

	if err != nil {
		fmt.Println("Could not read API")
		fmt.Println("Error: ", err.Error())
	}

	// Convert API data
	json.Unmarshal(groupieData, &data)
}
