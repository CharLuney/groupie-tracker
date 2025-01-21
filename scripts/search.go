package api

import (
	"regexp"
)

func Search(input string) {
	for i := 0; i < len(artists); i++ {
		query.isName, _ = regexp.MatchString(regex.Name, input)
		query.isDate, _ = regexp.MatchString(regex.Date, input)
		query.isYear, _ = regexp.MatchString(regex.Year, input)

		if query.isName {
		} else if query.isDate {

		} else if query.isYear {

		} else {
			query.isInvalid = true
		}
	}

}
