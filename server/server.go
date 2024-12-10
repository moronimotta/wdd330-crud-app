package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initServer() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	router.Run(":35536")
}
