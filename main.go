package main

import (
	"CrudApi/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	var book = models.Book{
		Title:  "libro",
		Author: "stuz zu",
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, book)

	})
	r.GET("/generate_uuid", func(c *gin.Context) {
		// Generate a new UUID
		newUUID := uuid.New()

		c.JSON(200, gin.H{
			"uuid": newUUID.String(),
		})
	})
	r.Run()
}
