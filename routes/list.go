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

type PostListRequest struct {
	ListTitle string `json:"title"`
	UserId    string `json:"user_id"`
}

func PostAnimeList(c *gin.Context) {
	var newAnimeList PostListRequest

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
	result := fmt.Sprintf("Next Anime List Created: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

func PostCharacterList(c *gin.Context) {
	var newCharacterList PostListRequest

	id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&newCharacterList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime list title"})
		return
	}
	newCharacterList.UserId = id.(string)

	insertedID, err := database.CreateCharacterList(newCharacterList.ListTitle, newCharacterList.UserId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Next Character List Created: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}
