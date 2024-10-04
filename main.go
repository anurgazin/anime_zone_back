package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "anime_zone/back_end/db"
	"anime_zone/back_end/middleware"
	"anime_zone/back_end/routes"
)

func getDefault(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello, World")
}

func main() {
	database.RunMongo()

	router := gin.Default()
	router.GET("/", getDefault)

	router.POST("/register", routes.Registration)
	router.POST("/login", routes.Login)
	router.PUT("/user/:id", middleware.AuthToken, routes.PutUser)

	router.GET("/anime", middleware.AuthToken, routes.GetAnime)
	router.GET("/anime/id/:id", routes.GetAnimeById)
	router.GET("/anime/title/:title", routes.GetAnimeByTitle)
	router.POST("/anime", middleware.AuthToken, routes.PostAnime)
	router.PUT("/anime/:id", middleware.AuthToken, routes.PutAnime)
	router.DELETE("/anime/:id", middleware.AuthToken, routes.DeleteAnime)

	router.GET("/characters", routes.GetCharacters)
	router.GET("/characters/:id", routes.GetCharactersById)
	router.POST("/characters", routes.PostCharacters)
	router.PUT("/characters/:id", routes.PutCharacters)
	router.DELETE("/characters/:id", routes.DeleteCharacter)

	router.Run("localhost:8080")
}
