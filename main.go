package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}

	dsn := "host=localhost user=postgres dbname=cosa sslmode=disable password=123"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to the database")
	db.AutoMigrate(&Product{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "hola mundo")
	})

	r.Run()

}
