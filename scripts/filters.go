package api

import (
	"strconv"
	"strings"
)

// Clear and apply filters to artists data
func ApplyFilters() {
	ClearFilters()
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
		}
	}
}

// Filters creation dates of each artist
func FilterCreation(i int) {
	date := strconv.Itoa(artists[i].CreationDate)
	if date != filters.CreationDate {
		filters.Verified = false
	}
}

// Filters first albums of each artist
func FilterAlbum(i int) {
	if !strings.Contains(filters.FirstAlbum, artists[i].FirstAlbum) {
		filters.Verified = false
	}
}

// Filters members of each artist
func FilterMembers(i int) {
	min, _ := strconv.Atoi(filters.MembersMin)
	max, _ := strconv.Atoi(filters.MembersMax)
	if !(len(artists[i].Members) >= min && len(artists[i].Members) <= max) {
		filters.Verified = false
	}
}

// Sets filters back to default
func ClearFilters() {
	artistsFilterted = artistsFilterted[:0]
}
