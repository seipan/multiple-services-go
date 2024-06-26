package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cat", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"animal": "cat"})
	})

	router.Run(":8081")
}
