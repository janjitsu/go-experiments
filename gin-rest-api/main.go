package main

import (
	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/albums", controllers.GetAlbums)

	router.GET("/albums/:id", controllers.GetAlbumByID)

	router.POST("/albums", controllers.PostAlbum)

	router.DELETE("/albums/:id", controllers.RemoveAlbumByID)

	router.Run("localhost:8080")
}
