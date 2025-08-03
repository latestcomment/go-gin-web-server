package main

import (
	"github.com/gin-gonic/gin"
	"github.com/latestcomment/go-gin-web-server/internal/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAllAlbums)
	router.GET("/album/:id", handlers.GetAlbumByID)

	router.Run("localhost:8081")
}
