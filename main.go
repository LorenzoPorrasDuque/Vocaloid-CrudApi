package main

import (
	"CrudApi/database"
	"CrudApi/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db := database.ConnectToDb()
	//create all tables
	db.AutoMigrate(models.Vocaloid{}, models.Song{})
	//insert data
	err := database.ExecuteSQLFile(db, "database/init.sql")
	if err != nil {
		log.Fatal(err)
	}
	//initialize the router
	r := gin.Default()

	//all the endpoints will be here

	//endpoint to databse all related

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
		case "songs":
			var song []models.Song
			db.Find(&song)
			data = song
		case "vocaloids":
			var vocaloids []models.Vocaloid
			db.Model(&vocaloids).Preload("Songs").Find(&vocaloids)
			data = vocaloids
		default:
			c.IndentedJSON(http.StatusNotFound, "Table not found")
			return
		}
		c.IndentedJSON(http.StatusOK, data)
	})

	//endpoints table vocaloids  ------------------------------------------------

	r.POST("/vocaloid", func(c *gin.Context) {
		var vocaloid models.Vocaloid
		if err := c.BindJSON(&vocaloid); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		db.Create(&vocaloid)
		c.IndentedJSON(http.StatusCreated, vocaloid)
	})
	r.GET("/vocaloid/:id", func(c *gin.Context) {
		var vocaloid models.Vocaloid
		id := c.Param("id")
		db.Model(&vocaloid).Preload("Songs").First(&vocaloid, id)
		c.IndentedJSON(http.StatusOK, vocaloid)
	})
	r.PUT("/vocaloid/:id", func(c *gin.Context) {
		var vocaloid models.Vocaloid
		id := c.Param("id")
		db.First(&vocaloid, id)
		if err := c.BindJSON(&vocaloid); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		db.Save(&vocaloid)
		c.IndentedJSON(http.StatusOK, vocaloid)
	})
	r.DELETE("/vocaloid/:id", func(c *gin.Context) {
		var vocaloid models.Vocaloid
		id := c.Param("id")
		db.First(&vocaloid, id)
		db.Delete(&vocaloid)
		c.IndentedJSON(http.StatusOK, "Vocaloid deleted")
	})

	//endpoints table Song ------------------------------------------------

	r.POST("/song", func(c *gin.Context) {
		var song models.Song
		if err := c.BindJSON(&song); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		db.Create(&song)
		c.IndentedJSON(http.StatusCreated, song)
	})
	r.GET("/song/:id", func(c *gin.Context) {
		var song models.Song
		id := c.Param("id")
		db.First(&song, id)
		c.IndentedJSON(http.StatusOK, song)
	})
	r.PUT("/song/:id", func(c *gin.Context) {
		var song models.Song
		id := c.Param("id")
		db.First(&song, id)
		if err := c.BindJSON(&song); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		db.Save(&song)
		c.IndentedJSON(http.StatusOK, song)
	})
	r.DELETE("/song/:id", func(c *gin.Context) {
		var song models.Song
		id := c.Param("id")
		db.First(&song, id)
		db.Delete(&song)
		c.IndentedJSON(http.StatusOK, "Song deleted")
	})

	r.Run()

}
