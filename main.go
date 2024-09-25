package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"anime_zone/back_end/funcs"
)

func getDefault(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello, World")
}

func main() {
	router := gin.Default()
	router.GET("/", getDefault)

	router.GET("/anime", funcs.GetAnime)
	router.GET("/anime/id/:id", funcs.GetAnimeById)
	router.GET("/anime/title/:title", funcs.GetAnimeByTitle)
	router.POST("/anime", funcs.PostAnime)
	router.PUT("/anime/:id", funcs.PutAnime)

	router.GET("/characters", funcs.GetCharacters)
	router.GET("/characters/:id", funcs.GetCharactersById)
	router.POST("/characters", funcs.PostCharacters)

	router.Run("localhost:8080")
}
