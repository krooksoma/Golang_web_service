package main

import(
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.POST("/albums", postAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8000")
}