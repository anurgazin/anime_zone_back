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
	client := database.RunMongo()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Auth", "Accept", "User-Agent", "Cache-Control", "Pragma", "RefreshToken"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.GET("/", getDefault)

	router.POST("/register", func(c *gin.Context) { routes.Registration(c, client) })
	router.POST("/login", func(c *gin.Context) { routes.Login(c, client) })
	router.GET("/user/:id", func(c *gin.Context) { routes.GetUser(c, client) })
	router.PUT("/user/:id", middleware.AuthToken, func(c *gin.Context) { routes.PutUser(c, client) })

	router.POST("/refresh", func(c *gin.Context) { middleware.RefreshToken(c, client) })

	router.GET("/anime", func(c *gin.Context) { routes.GetAnime(c, client) })
	router.GET("/anime/highest", func(c *gin.Context) { routes.GetHighestRatedAnime(c, client) })
	router.GET("/anime/popular", func(c *gin.Context) { routes.GetMostPopularAnime(c, client) })
	router.GET("/anime/similar/:id", func(c *gin.Context) { routes.GetSimilarAnime(c, client) })
	router.GET("/anime/id/:id", func(c *gin.Context) { routes.GetAnimeById(c, client) })

	router.GET("/anime/details/:id", func(c *gin.Context) { routes.GetAnimeDetails(c, client) })

	router.GET("/anime/list/:id", func(c *gin.Context) { routes.GetAllAnimeFromList(c, client) })

	router.GET("/anime/title/:title", func(c *gin.Context) { routes.GetAnimeByTitle(c, client) })
	router.GET("/anime/rating/:id", func(c *gin.Context) { routes.GetAnimeRatingById(c, client) })
	router.GET("/anime/rating/user/:id", func(c *gin.Context) { routes.GetAnimeRatingByUser(c, client) })
	router.POST("/anime", middleware.AuthToken, middleware.IsAdmin, func(c *gin.Context) { routes.PostAnime(c, client) })
	router.POST("/anime/rating/:id", middleware.AuthToken, func(c *gin.Context) { routes.RateAnime(c, client) })
	router.PUT("/anime/:id", middleware.AuthToken, middleware.IsAdmin, func(c *gin.Context) { routes.PutAnime(c, client) })
	router.DELETE("/anime/:id", middleware.AuthToken, middleware.IsAdmin, func(c *gin.Context) { routes.DeleteAnime(c, client) })

	router.GET("/characters", func(c *gin.Context) { routes.GetCharacters(c, client) })
	router.GET("/characters/name/asc", func(c *gin.Context) { routes.GetCharactersFirstName(c, client) })
	router.GET("/characters/id/:id", func(c *gin.Context) { routes.GetCharactersById(c, client) })
	router.GET("/characters/details/:id", func(c *gin.Context) { routes.GetCharacterDetails(c, client) })

	router.GET("/characters/list/:id", func(c *gin.Context) { routes.GetAllCharactersFromList(c, client) })

	router.GET("/characters/anime/:id", func(c *gin.Context) { routes.GetCharactersByAnimeId(c, client) })
	router.POST("/characters", middleware.AuthToken, middleware.IsAdmin, func(c *gin.Context) { routes.PostCharacters(c, client) })
	router.PUT("/characters/:id", middleware.AuthToken, middleware.IsAdmin, func(c *gin.Context) { routes.PutCharacters(c, client) })
	router.DELETE("/characters/:id", middleware.AuthToken, middleware.IsAdmin, func(c *gin.Context) { routes.DeleteCharacter(c, client) })

	router.GET("/list/anime", func(c *gin.Context) { routes.GetAnimeLists(c, client) })
	router.GET("/list/characters", func(c *gin.Context) { routes.GetCharacterLists(c, client) })
	router.GET("/list/anime/:id", func(c *gin.Context) { routes.GetAnimeListById(c, client) })
	router.GET("/list/characters/:id", func(c *gin.Context) { routes.GetCharacterListById(c, client) })
	router.GET("/list/anime/anime/:id", func(c *gin.Context) { routes.GetAnimeListsByAnimeId(c, client) })
	router.GET("/list/characters/character/:id", func(c *gin.Context) { routes.GetCharacterListsByCharacterId(c, client) })
	router.GET("/list/anime/user/:id", func(c *gin.Context) { routes.GetAnimeListsByUserId(c, client) })
	router.GET("/list/characters/user/:id", func(c *gin.Context) { routes.GetCharacterListsByUserId(c, client) })
	router.POST("/list/anime", middleware.AuthToken, func(c *gin.Context) { routes.PostAnimeList(c, client) })
	router.POST("/list/characters", middleware.AuthToken, func(c *gin.Context) { routes.PostCharacterList(c, client) })
	router.POST("/list/anime/add/:id", middleware.AuthToken, func(c *gin.Context) { routes.AddAnimeToList(c, client) })
	router.POST("/list/characters/add/:id", middleware.AuthToken, func(c *gin.Context) { routes.AddCharacterToList(c, client) })
	router.PATCH("/list/anime/edit/:id", middleware.AuthToken, func(c *gin.Context) { routes.EditAnimeList(c, client) })
	router.PATCH("/list/characters/edit/:id", middleware.AuthToken, func(c *gin.Context) { routes.EditCharacterList(c, client) })
	router.PATCH("/list/anime/rating/:id", middleware.AuthToken, func(c *gin.Context) { routes.UpdateAnimeListRating(c, client) })
	router.PATCH("/list/characters/rating/:id", middleware.AuthToken, func(c *gin.Context) { routes.UpdateCharacterListRating(c, client) })

	router.GET("/comment", func(c *gin.Context) { routes.GetAllComments(c, client) })
	router.GET("/comment/type/:type", func(c *gin.Context) { routes.GetCommentByType(c, client) })
	router.GET("/comment/id/:id", func(c *gin.Context) { routes.GetCommentById(c, client) })
	router.GET("/comment/:type/:id", func(c *gin.Context) { routes.GetCommentForContent(c, client) })
	router.GET("/comment/user/:id", func(c *gin.Context) { routes.GetCommentForUser(c, client) })
	router.POST("/comment", middleware.AuthToken, func(c *gin.Context) { routes.PostComment(c, client) })
	router.PATCH("/comment/id/:id", middleware.AuthToken, func(c *gin.Context) { routes.UpdateComment(c, client) })
	router.PATCH("/comment/rating/:id", middleware.AuthToken, func(c *gin.Context) { routes.UpdateCommentRating(c, client) })
	router.DELETE("/comment/id/:id", middleware.AuthToken, func(c *gin.Context) { routes.DeleteComment(c, client) })

	router.ServeHTTP(w, r)
}
