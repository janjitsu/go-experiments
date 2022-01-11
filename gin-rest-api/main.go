package main

import (
	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/handlers"

	"github.com/janjitsu/go-experiments/gin-rest-api/platform/albums"
)

func main() {

	albumRepo := albums.New()

	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums(albumRepo))

	router.GET("/albums/:id", handlers.GetAlbumByID(albumRepo))

	router.POST("/albums", handlers.PostAlbum(albumRepo))

	router.DELETE("/albums/:id", handlers.RemoveAlbumByID(albumRepo))

	router.Run("localhost:8080")
}
