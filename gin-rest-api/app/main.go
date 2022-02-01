package main

import (
	"github.com/gin-gonic/gin"

	"github.com/janjitsu/go-experiments/gin-rest-api/app/handlers"

	"github.com/janjitsu/go-experiments/gin-rest-api/app/config"

	"github.com/janjitsu/go-experiments/gin-rest-api/app/platform/albums"
)

func main() {

	var configuration = config.SetupConfig("../.")

	albumRepo := albums.NewDbRepo(configuration.Database.ConnectionUri)

	//albumRepo := albums.New()

	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums(albumRepo))

	router.GET("/albums/:id", handlers.GetAlbumByID(albumRepo))

	router.POST("/albums", handlers.PostAlbum(albumRepo))

	router.DELETE("/albums/:id", handlers.RemoveAlbumByID(albumRepo))

	router.Run("localhost:" + configuration.Server.Port)
}
