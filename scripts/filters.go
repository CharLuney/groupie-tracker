package api

import (
	"fmt"
	"strconv"
	"strings"
)

func ApplyFilters() {
	fmt.Println("Clearing filters..")
	ClearFilters()
	fmt.Println("Filtering artists..")
	for i := 0; i < len(artists); i++ {
		filters.Verified = true
		if filters.CreationDate != "" {
			FilterCreation(i)
		}
		if filters.FirstAlbum != "" {
			FilterAlbum(i)
		}
		if filters.MembersMax != "" && filters.MembersMin != "" {
			FilterMembers(i)
		}

		if filters.Verified {
			artistsFilterted = append(artistsFilterted, artists[i])
			fmt.Println("Added artist: ", artists[i])
		}
	}
}

func FilterCreation(i int) {
	date := strconv.Itoa(artists[i].CreationDate)
	if date != filters.CreationDate {
		filters.Verified = false
	}
}

func FilterAlbum(i int) {
	if !strings.Contains(filters.FirstAlbum, artists[i].FirstAlbum) {
		filters.Verified = false
	}
}

func FilterMembers(i int) {
	min, _ := strconv.Atoi(filters.MembersMin)
	max, _ := strconv.Atoi(filters.MembersMax)
	if !(len(artists[i].Members) >= min && len(artists[i].Members) <= max) {
		filters.Verified = false
	}
}

func FilterLocations(i int) {

}

func ClearFilters() {
	artistsFilterted = artistsFilterted[:0]
}
