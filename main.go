package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/krooksoma/Golang-web-service/domain"
)


//get endpoint
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

//post endpoint
func postAlbums(c *gin.Context){
	var newAlbum album

	//Call BindJSON to bind the receiver JSON
	//to the newAlbum

	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}

	//Add the new Album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//get /albums/:id

func getAlbumById(c *gin.Context){
	id := c.Param("id")

	//Loop over list of albuns, looking for
	// an album whose ID value matches the parameter

	for _, a := range albums{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
	router := gin.Default()
	router.POST("/albums", postAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8000")
}