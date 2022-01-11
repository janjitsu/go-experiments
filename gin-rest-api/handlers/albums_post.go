package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/platform/albums"
)

func PostAlbum(albumRepo albums.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newAlbum albums.Album

		if err := c.BindJSON(&newAlbum); err != nil { //default error checking
			//&newAlbum is a pointer to the memory address of the variable
			return
		}

		albumRepo.Add(newAlbum)
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}
}
