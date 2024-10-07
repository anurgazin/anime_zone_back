package routes

import (
	database "anime_zone/back_end/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

type AddToListRequest struct {
	ListID   string `json:"list_id"`
	UserID   string `json:"user_id"`
	ObjectID string `json:"object_id"`
}

func AddAnimeToList(c *gin.Context) {
	var newAnimeList AddToListRequest

	uID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	lID := c.Param("id")
	if lID == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "List Id not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&newAnimeList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime list title"})
		return
	}
	newAnimeList.UserID = uID.(string)
	newAnimeList.ListID = lID

	result, err := database.AddAnimeToList(newAnimeList.ListID, newAnimeList.UserID, newAnimeList.ObjectID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}

func AddCharacterToList(c *gin.Context) {
	var newCharacterList AddToListRequest

	uID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	lID := c.Param("id")
	if lID == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "List Id not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&newCharacterList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime list title"})
		return
	}
	newCharacterList.UserID = uID.(string)
	newCharacterList.ListID = lID

	result, err := database.AddCharacterToList(newCharacterList.ListID, newCharacterList.UserID, newCharacterList.ObjectID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}

func GetAnimeLists(c *gin.Context) {
	animeList, err := database.GetAllAnimeLists()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime lists"})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterLists(c *gin.Context) {
	characterList, err := database.GetAllCharacterLists()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve character lists"})
		return
	}
	c.IndentedJSON(http.StatusOK, characterList)
}

func GetAnimeListById(c *gin.Context) {
	id := c.Param("id")

	anime, err := database.GetAnimeListById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}
