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
	router.GET("/anime/:id", funcs.GetAnimeById)
	router.POST("/anime", funcs.PostAnime)

	router.Run("localhost:8080")
}
