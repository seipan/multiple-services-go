package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/animal", func(c *gin.Context) {
		resp, err := http.Get("http://0.0.0.0:8081/cat")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cat"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		c.Data(http.StatusOK, "application/json", body)
	})

	router.Run(":8082")
}
