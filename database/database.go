package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToDb() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env files")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	fmt.Println("Connected to the database")
	return database
}

func ExecuteSQLFile(db *gorm.DB, filePath string) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Convert []byte to string
	sqlScript := string(fileContent)
	fmt.Print(sqlScript)

	// Execute SQL script
	result := db.Exec(sqlScript)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
