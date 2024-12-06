package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type Artist struct {
	ID   string `json:id`
	Name string `json:"name"`
}

var Artists = []Artist{
	{ID: "1", Name: "Queen"},
	{ID: "2", Name: "SOJA"},
	{ID: "3", Name: "Pink Floyd"},
}

func Search(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Artists)
}

func main() {
	router := gin.Default()
	router.GET("/Artists", Search)
	router.Run("localhost:8080")
}
