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
	//preguntarle a a sebas diferencia entre & y * en este caso, para saber si se ahce copia o referencia
	db.AutoMigrate(models.Book{}, models.Author{})
	db.Migrator().CurrentDatabase()
	db.Migrator().GetTables()

	//initialize the router
	r := gin.Default()

	//all the endpoints will be here
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

	r.Run()

}
