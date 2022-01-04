package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Struct is how we define and organize data to be passed around, similar to a data class
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// We define a list of structs
var albums = []album{
	{ID: "1", Title: "Fortunate", Artist: "Gojira", Price: 99.99},
	{ID: "2", Title: "Amazonia", Artist: "Gojira", Price: 99.99},
	{ID: "3", Title: "Hold On", Artist: "Gojira", Price: 99.99},
	{ID: "4", Title: "New Found", Artist: "Gojira", Price: 99.99},
}

func getAlbums(c *gin.Context) { //pointer to variable value
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil { //default error checking
		//&newAlbum is a pointer to the memory address of the variable
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func removeAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:len(albums)]...) //working with slices
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbum)

	router.DELETE("/albums/:id", removeAlbumByID)

	router.Run("localhost:8080")
}
