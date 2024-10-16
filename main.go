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

	router.GET("/anime", routes.GetAnime)
	router.GET("/anime/id/:id", routes.GetAnimeById)
	router.GET("/anime/title/:title", routes.GetAnimeByTitle)
	router.POST("/anime", middleware.AuthToken, middleware.IsAdmin, routes.PostAnime)
	router.PUT("/anime/:id", middleware.AuthToken, middleware.IsAdmin, routes.PutAnime)
	router.DELETE("/anime/:id", middleware.AuthToken, middleware.IsAdmin, routes.DeleteAnime)

	router.GET("/characters", routes.GetCharacters)
	router.GET("/characters/:id", routes.GetCharactersById)
	router.POST("/characters", middleware.AuthToken, middleware.IsAdmin, routes.PostCharacters)
	router.PUT("/characters/:id", middleware.AuthToken, middleware.IsAdmin, routes.PutCharacters)
	router.DELETE("/characters/:id", middleware.AuthToken, middleware.IsAdmin, routes.DeleteCharacter)

	router.GET("/list/anime", routes.GetAnimeLists)
	router.GET("/list/characters", routes.GetCharacterLists)
	router.GET("/list/anime/:id", routes.GetAnimeListById)
	router.GET("/list/characters/:id", routes.GetCharacterListById)
	router.POST("/list/anime", middleware.AuthToken, routes.PostAnimeList)
	router.POST("/list/characters", middleware.AuthToken, routes.PostCharacterList)
	router.PATCH("/list/anime/add/:id", middleware.AuthToken, routes.AddAnimeToList)
	router.PATCH("/list/characters/add/:id", middleware.AuthToken, routes.AddCharacterToList)
	router.PATCH("/list/anime/edit/:id", middleware.AuthToken, routes.EditAnimeList)
	router.PATCH("/list/characters/edit/:id", middleware.AuthToken, routes.EditCharacterList)

	router.GET("/comment", routes.GetAllComments)
	router.GET("/comment/type/:type", routes.GetCommentByType)
	router.GET("/comment/id/:id", routes.GetCommentById)
	router.POST("/comment", middleware.AuthToken, routes.PostComment)
	router.PATCH("/comment/id/:id", middleware.AuthToken, routes.UpdateComment)
	router.DELETE("/comment/id/:id", middleware.AuthToken, routes.DeleteComment)

	router.Run("localhost:8080")
}
