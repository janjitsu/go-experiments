package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/models"
)

func GetAlbums(c *gin.Context) { //pointer to variable value
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil { //default error checking
		//&newAlbum is a pointer to the memory address of the variable
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, models.Albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func RemoveAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range models.Albums {
		if a.ID == id {
			models.Albums = append(models.Albums[:i], models.Albums[i+1:len(models.Albums)]...) //working with slices
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
