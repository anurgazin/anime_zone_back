package routes

import (
	database "anime_zone/back_end/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAnimeList(c *gin.Context) {
	var newAnimeList database.PostListRequest

	id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	username, exists := c.Get("username")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&newAnimeList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime list title"})
		return
	}
	newAnimeList.UserId = id.(string)
	newAnimeList.Username = username.(string)

	insertedID, err := database.CreateAnimeList(newAnimeList)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Next Anime List Created: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

func PostCharacterList(c *gin.Context) {
	var newCharacterList database.PostListRequest

	id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	username, exists := c.Get("username")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&newCharacterList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime list title"})
		return
	}
	newCharacterList.UserId = id.(string)
	newCharacterList.Username = username.(string)

	insertedID, err := database.CreateCharacterList(newCharacterList)
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

	animeList, err := database.GetAnimeListById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterListById(c *gin.Context) {
	id := c.Param("id")

	characterList, err := database.GetCharacterListById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, characterList)
}

type UpdateList struct {
	Name string `bson:"name" json:"name"`
}

func EditAnimeList(c *gin.Context) {
	id := c.Param("id")
	var updateAnimeList UpdateList

	user_id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&updateAnimeList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text"})
		return
	}

	result, err := database.UpdateAnimeList(id, user_id.(string), updateAnimeList.Name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func EditCharacterList(c *gin.Context) {
	id := c.Param("id")
	var updateCharList UpdateList

	user_id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&updateCharList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text"})
		return
	}

	result, err := database.UpdateCharacterList(id, user_id.(string), updateCharList.Name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

type UpdateListRatingAction struct {
	Action string `json:"action" form:"action"` // "increment" or "decrement"
}

func UpdateAnimeListRating(c *gin.Context) {
	id := c.Param("id")
	var updateData UpdateListRatingAction

	// Bind the action (increment or decrement) from the request
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Determine the increment value based on the action
	var incrementValue float64
	if updateData.Action == "increment" {
		incrementValue = 1.0
	} else if updateData.Action == "decrement" {
		incrementValue = -1.0
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid action. Use 'increment' or 'decrement'."})
		return
	}

	_, err := database.UpdateAnimeListRating(id, incrementValue)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update anime list rating or anime list not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Anime List rating updated successfully"})
}

func UpdateCharacterListRating(c *gin.Context) {
	id := c.Param("id")
	var updateData UpdateListRatingAction

	// Bind the action (increment or decrement) from the request
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Determine the increment value based on the action
	var incrementValue float64
	if updateData.Action == "increment" {
		incrementValue = 1.0
	} else if updateData.Action == "decrement" {
		incrementValue = -1.0
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid action. Use 'increment' or 'decrement'."})
		return
	}

	_, err := database.UpdateCharacterListRating(id, incrementValue)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update character list rating or character list not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Character List rating updated successfully"})
}

func GetAnimeListsByAnimeId(c *gin.Context) {
	id := c.Param("id")

	animeList, err := database.GetAllAnimeListsByAnimeId(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterListsByCharacterId(c *gin.Context) {
	id := c.Param("id")

	animeList, err := database.GetAllCharacterListsByCharacterId(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetAnimeListsByUserId(c *gin.Context) {
	id := c.Param("id")

	animeList, err := database.GetAllAnimeListsByUserId(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterListsByUserId(c *gin.Context) {
	id := c.Param("id")

	animeList, err := database.GetAllCharacterListsByUserId(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}
