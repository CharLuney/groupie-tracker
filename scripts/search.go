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
	fmt.Println("regex string: ", query.IsName)
	query.IsDate, _ = regexp.MatchString(regex.Date, input)
	fmt.Println("regex date: ", query.IsDate)
	query.IsYear, _ = regexp.MatchString(regex.Year, input)
	fmt.Println("regex year: ", query.IsYear)

	// Looping through artists
	for i := 0; i < len(artists); i++ {
		query.LookedFor = false
		query.IsInvalid = false

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
