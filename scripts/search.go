package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SearchFor(input string) {

	ClearSearch()
	input = strings.ToLower(input)

	// Defines input regex
	query.IsName, _ = regexp.MatchString(regex.Name, input)
	query.IsDate, _ = regexp.MatchString(regex.Date, input)
	query.IsYear, _ = regexp.MatchString(regex.Year, input)

	// Looping through artists
	for i := 0; i < len(artists); i++ {
		query.LookedFor = false
		query.IsInvalid = false

		if query.IsName {
			if strings.Contains(strings.ToLower(artists[i].Name), input) {
				query.LookedFor = true
			}
			for j := 0; j < len(artists[i].Members); j++ {
				if strings.Contains(strings.ToLower(artists[i].Members[j]), input) {
					query.LookedFor = true
				}
			}
		} else if query.IsDate {
			if strings.Contains(artists[i].FirstAlbum, input) {
				query.LookedFor = true
			}

		} else if query.IsYear {
			inputInt, _ := strconv.Atoi(input)
			if artists[i].CreationDate == inputInt {
				query.LookedFor = true
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

func ClearSearch() {
	artistsSearched = artistsSearched[:0]
}
