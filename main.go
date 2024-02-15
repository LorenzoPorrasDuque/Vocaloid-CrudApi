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

	//initialize the router
	r := gin.Default()

	//all the endpoints will be here
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "hola mundo")
	})

	r.Run()

}
