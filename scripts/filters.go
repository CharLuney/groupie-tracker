package api

import (
	"fmt"
	"strconv"
)

func ApplyFilters() {
	fmt.Println("Clearing filters..")
	ClearFilters()
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
	date, _ := strconv.Atoi(filters.CreationDate)
	if date != artists[i].CreationDate {
		filters.Satisfied = false
	}
}

func FilterAlbum(i int) {
	// albumYear := strconv.Itoa(artists[i].FirstAlbum)
	// albumYear = albumYear[len(albumYear)-4:]
	// date, _ := strconv.Atoi(albumYear)
	// fmt.Println(date)
	// if date != {
	// 	// strconv.Atoi(artists[i].FirstAlbum[len(artists[i].FirstAlbum)-4:]

	// 	filters.Satisfied = false
	// }
}

func FilterMembers(i int) {
	amt, _ := strconv.Atoi(filters.MembersAmount)
	if len(artists[i].Members) != amt {
		filters.Satisfied = false
	}
}

func FilterLocations(i int) {

}

func ClearFilters() {
	artistsFilterted = artistsFilterted[:0]
}
