package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetMap(URL string) []byte {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	resp.Body.Close()
	return body
}
