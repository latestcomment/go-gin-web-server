package services

import (
	"github.com/latestcomment/go-gin-web-server/internal/models"
	"github.com/latestcomment/go-gin-web-server/internal/repositories"
)

func GetAllAlbums() []models.Album {
	return repositories.GetAllAlbums()
}

func GetAlbumByID(id string) (models.Album, bool) {
	return repositories.GetAlbumByID(id)
}
