package main

import (
	"CrudApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println("Connected to the database")
	db.AutoMigrate(&models.Book{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "hola mundo")
	})

	r.Run()

}
