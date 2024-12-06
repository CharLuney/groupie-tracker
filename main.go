package main

import (
	"github.com/gin-gonic/gin"
	//"errors"
)

func Search(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, artists)
}

func main() {
	router := gin.Default()
	router.GET("/artists", Search)
	router.Run("localhost:8080")
}
