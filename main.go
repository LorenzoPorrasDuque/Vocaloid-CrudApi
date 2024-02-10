package main

import (
	"CrudApi/models"
	"github.com/gin-gonic/gin"
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
	r.Run()
}
