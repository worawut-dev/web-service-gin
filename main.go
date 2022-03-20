package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ablum struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []ablum{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum ablum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	paramID := c.Param("id")
	for _, data := range albums {
		if data.ID == paramID {
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}

func updateAlbums(c *gin.Context) {
	var editAlbum ablum

	if err := c.BindJSON(&editAlbum); err != nil {
		return
	}

	paramID := c.Param("id")
	for i := 0; i <= len(albums)-1; i++ {
		if albums[i].ID == paramID {
			albums[i].ID = editAlbum.ID
			albums[i].Title = editAlbum.Title
			albums[i].Artist = editAlbum.Artist
			albums[i].Price = editAlbum.Price

			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, "data not found")
}

func deleteAlbumById(c *gin.Context) {
	paramID := c.Param("id")
	for i := 0; i <= len(albums)-1; i++ {
		if albums[i].ID == paramID {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, "delete success")
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("page/*.html")
	router.GET("/", getHome)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbums)
	router.DELETE("/albums/:id", deleteAlbumById)

	router.Run("localhost:8080")

}
