package api

import (
	"fmt"
	"regexp"
)

func Search() {
	var searched = "test1"
	var isMatching, _ = regexp.MatchString("[a-z]+$", searched)
	fmt.Println(isMatching)
}
