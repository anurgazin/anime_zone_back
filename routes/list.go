package routes

import (
	database "anime_zone/back_end/db"
	// "anime_zone/back_end/funcs"
	"fmt"

	// "fmt"
	// "io"
	"net/http"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type PostAnimeListRequest struct {
	ListTitle string `json:"title"`
	UserId    string `json:"user_id"`
}

func PostAnimeList(c *gin.Context) {
	var newAnimeList PostAnimeListRequest

	id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&newAnimeList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime list title"})
		return
	}
	newAnimeList.UserId = id.(string)

	insertedID, err := database.CreateAnimeList(newAnimeList.ListTitle, newAnimeList.UserId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Next List Created: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}
