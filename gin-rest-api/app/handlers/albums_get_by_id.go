package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/app/platform/albums"
)

func GetAlbumByID(albumRepo albums.GetterById) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		album, err := albumRepo.GetById(id)

		if album != nil {
			c.IndentedJSON(http.StatusOK, album)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
}
