package api

import (
	"fmt"
	"regexp"
)

func Search() {
	var searched = "test1."
	var isMatching, err = regexp.MatchString(`[a-z]+`, searched)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if isMatching {
		fmt.Println(isMatching)
	}
}
