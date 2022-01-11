package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/platform/albums"
)

func GetAlbums(albumRepo albums.Getter) gin.HandlerFunc { //pointer to variable value
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albumRepo.GetAll())
	}
}
