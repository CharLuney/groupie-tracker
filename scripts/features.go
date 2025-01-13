package api

import (
	"fmt"
	"strconv"
)

func Search() {

}

func FilterCreation() {

}

func FilterAlbum() {

}

func FilterMembers() {
	fmt.Println("Filtering artists..")
	for i := 0; i < len(artists); i++ {
		amt, _ := strconv.Atoi(filter.MembersAmount)
		if len(artists[i].Members) == amt {
			artistsFilterted = append(artistsFilterted, artists[i])
			fmt.Println("Added artist: ", artists[i])
		}
	}
}

func FilterLocations() {

}
