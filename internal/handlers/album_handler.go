package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/latestcomment/go-gin-web-server/internal/services"
)

func GetAllAlbums(c *gin.Context) {
	albums := services.GetAllAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, ok := services.GetAlbumByID(id)
	if ok {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
