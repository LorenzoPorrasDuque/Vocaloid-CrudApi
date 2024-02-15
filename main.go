package main

import (
	"CrudApi/database"
	"CrudApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db := database.ConnectToDb()
	//create all tables
	db.AutoMigrate(models.Book{}, models.Author{})
	db.Migrator().CurrentDatabase()
	db.Migrator().GetTables()

	//initialize the router
	r := gin.Default()

	//all the endpoints will be here

	//enpoint to databse all related
	r.GET("/getDatabase", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, db.Migrator().CurrentDatabase())
	})

	r.GET("/getTables", func(c *gin.Context) {
		tables, err := db.Migrator().GetTables()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Respond with the list of tables as JSON
		c.IndentedJSON(http.StatusOK, tables)
	})

	//get data from tables
	r.GET("/data/:table", func(c *gin.Context) {
		table := c.Param("table")
		var data interface{}
		switch table {
		case "books":
			var books []models.Book
			db.Find(&books)
			data = books
		case "authors":
			var authors []models.Author
			db.Find(&authors)
			data = authors
		default:
			c.IndentedJSON(http.StatusNotFound, "Table not found")
			return
		}
		c.IndentedJSON(http.StatusOK, data)
	})
	//endpoints table books
	r.POST("/books", func(c *gin.Context) {
		var book models.Book
		if err := c.BindJSON(&book); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		db.Create(&book)
		c.IndentedJSON(http.StatusCreated, book)
	})

	r.GET("/books/:id", func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		db.First(&book, id)
		c.IndentedJSON(http.StatusOK, book)
	})

	r.PUT("/books/:id", func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		db.First(&book, id)
		if err := c.BindJSON(&book); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		db.Save(&book)
		c.IndentedJSON(http.StatusOK, book)
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		db.First(&book, id)
		db.Delete(&book)
		c.IndentedJSON(http.StatusOK, "Book deleted")
	})

	//endpoints table authors
	
	r.Run()

}
