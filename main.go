package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json: "id"`
	Title  string `json: "title"`
	Artist string `json: artist`
	Price  string `json: price`
}

var albums = []album{
	{
		ID:     "1",
		Title:  "Zeit",
		Artist: "Rammstein",
		Price:  "40$",
	},
	{
		ID:     "2",
		Title:  "While we sleep",
		Artist: "Insomnium",
		Price:  "40$",
	},
	{
		ID:     "3",
		Title:  "Peninsula",
		Artist: "Dinosaur Pile-up",
		Price:  "40$",
	},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.Run("localhost:8000")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
