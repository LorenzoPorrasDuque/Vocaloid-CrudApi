package main

import (
	"CrudApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	var _ = models.Book{
		Title:  "libro",
		Author: "stuz zu",
	}
	dsn := "host=localhost user=postgres dbname=test sslmode=disable password=mysecretpassword"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to the database")
	db.AutoMigrate(&models.Book{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, book)

	})

	r.Run()

}
