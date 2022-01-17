package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/app/platform/albums"
)

func RemoveAlbumByID(albumRepo albums.Remover) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := albumRepo.RemoveById(id)
		if err == nil {
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
}
