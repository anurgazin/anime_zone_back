package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	database "anime_zone/back_end/db"
	"anime_zone/back_end/middleware"
	"anime_zone/back_end/routes"
)

func getDefault(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello, World")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	database.RunMongo()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Auth", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.GET("/", getDefault)

	router.POST("/register", routes.Registration)
	router.POST("/login", routes.Login)
	router.GET("/user/:id", routes.GetUser)
	router.PUT("/user/:id", middleware.AuthToken, routes.PutUser)

	router.GET("/anime", routes.GetAnime)
	router.GET("/anime/highest", routes.GetHighestRatedAnime)
	router.GET("/anime/popular", routes.GetMostPopularAnime)
	router.GET("/anime/id/:id", routes.GetAnimeById)
	router.GET("/anime/title/:title", routes.GetAnimeByTitle)
	router.GET("/anime/rating/:id", routes.GetAnimeRatingById)
	router.GET("/anime/rating/user/:id", routes.GetAnimeRatingByUser)
	router.POST("/anime", middleware.AuthToken, middleware.IsAdmin, routes.PostAnime)
	router.POST("/anime/rating/:id", middleware.AuthToken, routes.RateAnime)
	router.PUT("/anime/:id", middleware.AuthToken, middleware.IsAdmin, routes.PutAnime)
	router.DELETE("/anime/:id", middleware.AuthToken, middleware.IsAdmin, routes.DeleteAnime)

	router.GET("/characters", routes.GetCharacters)
	router.GET("/characters/name/asc", routes.GetCharactersFirstName)
	router.GET("/characters/id/:id", routes.GetCharactersById)
	router.GET("/characters/anime/:id", routes.GetCharactersByAnimeId)
	router.POST("/characters", middleware.AuthToken, middleware.IsAdmin, routes.PostCharacters)
	router.PUT("/characters/:id", middleware.AuthToken, middleware.IsAdmin, routes.PutCharacters)
	router.DELETE("/characters/:id", middleware.AuthToken, middleware.IsAdmin, routes.DeleteCharacter)

	router.GET("/list/anime", routes.GetAnimeLists)
	router.GET("/list/characters", routes.GetCharacterLists)
	router.GET("/list/anime/:id", routes.GetAnimeListById)
	router.GET("/list/characters/:id", routes.GetCharacterListById)
	router.GET("/list/anime/anime/:id", routes.GetAnimeListsByAnimeId)
	router.GET("/list/characters/character/:id", routes.GetCharacterListsByCharacterId)
	router.GET("/list/anime/user/:id", routes.GetAnimeListsByUserId)
	router.GET("/list/characters/user/:id", routes.GetCharacterListsByUserId)
	router.POST("/list/anime", middleware.AuthToken, routes.PostAnimeList)
	router.POST("/list/characters", middleware.AuthToken, routes.PostCharacterList)
	router.POST("/list/anime/add/:id", middleware.AuthToken, routes.AddAnimeToList)
	router.POST("/list/characters/add/:id", middleware.AuthToken, routes.AddCharacterToList)
	router.PATCH("/list/anime/edit/:id", middleware.AuthToken, routes.EditAnimeList)
	router.PATCH("/list/characters/edit/:id", middleware.AuthToken, routes.EditCharacterList)
	router.PATCH("/list/anime/rating/:id", middleware.AuthToken, routes.UpdateAnimeListRating)
	router.PATCH("/list/characters/rating/:id", middleware.AuthToken, routes.UpdateCharacterListRating)

	router.GET("/comment", routes.GetAllComments)
	router.GET("/comment/type/:type", routes.GetCommentByType)
	router.GET("/comment/id/:id", routes.GetCommentById)
	router.GET("/comment/:type/:id", routes.GetCommentForContent)
	router.GET("/comment/user/:id", routes.GetCommentForUser)
	router.POST("/comment", middleware.AuthToken, routes.PostComment)
	router.PATCH("/comment/id/:id", middleware.AuthToken, routes.UpdateComment)
	router.PATCH("/comment/rating/:id", middleware.AuthToken, routes.UpdateCommentRating)
	router.DELETE("/comment/id/:id", middleware.AuthToken, routes.DeleteComment)

	router.ServeHTTP(w, r)
}
