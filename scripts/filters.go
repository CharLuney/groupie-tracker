package api

import (
	"fmt"
	"strconv"
)

func ApplyFilters() {
	fmt.Println("Filtering artists..")
	for i := 0; i < len(artists); i++ {
		filters.Satisfied = true
		if filters.CreationDate != "" {
			FilterCreation(i)
		}
		if filters.FirstAlbum != "" {
			FilterAlbum(i)
		}
		if filters.MembersAmount != "" {
			FilterMembers(i)
		}
		if filters.ConcertLocation != "" {
			FilterLocations(i)
		}
		if filters.Satisfied {
			artistsFilterted = append(artistsFilterted, artists[i])
			fmt.Println("Added artist: ", artists[i])
		}
	}
	fmt.Println("Artists filtered, you can now apply")
}

func FilterCreation(i int) {

}

func FilterAlbum(i int) {

}

func FilterMembers(i int) {
	amt, _ := strconv.Atoi(filters.MembersAmount)
	if len(artists[i].Members) != amt {
		filters.Satisfied = false
	}
}

func FilterLocations(i int) {

}
