package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Search(input string) {
	// Looping through artists
	for i := 0; i < len(artists); i++ {
		query.LookedFor = false

		// Defines input regex
		query.IsName, _ = regexp.MatchString(regex.Name, input)
		query.IsDate, _ = regexp.MatchString(regex.Date, input)
		query.IsYear, _ = regexp.MatchString(regex.Year, input)

		if query.IsName {
			if strings.Contains(artists[i].Name, input) {
				query.LookedFor = true
			}
			for i := 0; i < len(artists[i].Members); i++ {
				if strings.Contains(artists[i].Members[i], input) {
					query.LookedFor = true
				}
			}
		} else if query.IsDate {
			if strings.Contains(artists[i].FirstAlbum, input) {
				query.LookedFor = true
			}

		} else if query.IsYear {
			inputInt = strconv.Atoi(input)
			if strings.Contains(artists[i].CreationDate, strconv.Atoi(input)) {

			}
		} else {
			query.IsInvalid = true
		}

		if !query.IsInvalid && query.LookedFor {
			artistsSearched = append(artistsSearched, artists[i])
			fmt.Println("Added artist to result: ", artists[i])
		}
	}

}
