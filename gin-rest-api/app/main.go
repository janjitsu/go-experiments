package main

import (
	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/app/handlers"

	"github.com/janjitsu/go-experiments/gin-rest-api/app/platform/albums"
)

func main() {

	albumRepo := albums.NewDbRepo("postgresql://gouser:123456@localhost:5432/gin-rest-api")

	//albumRepo := albums.New()

	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums(albumRepo))

	router.GET("/albums/:id", handlers.GetAlbumByID(albumRepo))

	router.POST("/albums", handlers.PostAlbum(albumRepo))

	router.DELETE("/albums/:id", handlers.RemoveAlbumByID(albumRepo))

	router.Run("localhost:8080")
}
