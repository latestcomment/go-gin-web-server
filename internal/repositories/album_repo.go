package repositories

import (
	"github.com/latestcomment/go-gin-web-server/internal/models"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 2010, Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 2011, Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 2012, Price: 39.99},
}

func GetAllAlbums() []models.Album {
	return albums
}

func GetAlbumByID(id string) (models.Album, bool) {
	for _, a := range albums {
		if a.ID == id {
			return a, true
		}
	}
	return models.Album{}, false
}
