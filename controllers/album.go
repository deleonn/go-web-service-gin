package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"web-service-gin/models"
)

type AlbumController struct{}

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Cliford Brown", Artist: "Sara Vaughan", Price: 39.99},
}

// GetAlbums response with the list of all albums as JSON.
func (a AlbumController) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds an album from JSON received in the request body.
func (a AlbumController) PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID retrieves an album by the id received in the request params.
func (a AlbumController) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
