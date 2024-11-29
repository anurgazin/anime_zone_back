package routes

import (
	database "anime_zone/back_end/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostAnimeList(c *gin.Context, client *mongo.Client) {
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

	insertedID, err := database.CreateAnimeList(newAnimeList, client)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Next Anime List Created: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

func PostCharacterList(c *gin.Context, client *mongo.Client) {
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

	insertedID, err := database.CreateCharacterList(newCharacterList, client)
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

func AddAnimeToList(c *gin.Context, client *mongo.Client) {
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

	result, err := database.AddAnimeToList(newAnimeList.ListID, newAnimeList.UserID, newAnimeList.ObjectID, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}

func AddCharacterToList(c *gin.Context, client *mongo.Client) {
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

	result, err := database.AddCharacterToList(newCharacterList.ListID, newCharacterList.UserID, newCharacterList.ObjectID, client)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}

func GetAnimeLists(c *gin.Context, client *mongo.Client) {
	animeList, err := database.GetAllAnimeLists(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime lists"})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterLists(c *gin.Context, client *mongo.Client) {
	characterList, err := database.GetAllCharacterLists(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve character lists"})
		return
	}
	c.IndentedJSON(http.StatusOK, characterList)
}

func GetAnimeListById(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	animeList, err := database.GetAnimeListById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterListById(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	characterList, err := database.GetCharacterListById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, characterList)
}

type UpdateList struct {
	Name string `bson:"name" json:"name"`
}

func EditAnimeList(c *gin.Context, client *mongo.Client) {
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

	result, err := database.UpdateAnimeList(id, user_id.(string), updateAnimeList.Name, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func EditCharacterList(c *gin.Context, client *mongo.Client) {
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

	result, err := database.UpdateCharacterList(id, user_id.(string), updateCharList.Name, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

type UpdateListRatingAction struct {
	ListType string `json:"list_type" form:"list_type"`
	Action   string `json:"action" form:"action"` // "increment" or "decrement"
}

func UpdateListRating(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	user_id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	username, exists := c.Get("username")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	var updateData UpdateListRatingAction

	// Bind the action (increment or decrement) from the request
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Determine the increment value based on the action
	var incrementValue int
	if updateData.Action == "increment" {
		incrementValue = 1
	} else if updateData.Action == "decrement" {
		incrementValue = -1
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid action. Use 'increment' or 'decrement'."})
		return
	}

	result, err := database.UpdateListRating(id, updateData.ListType, user_id.(string), username.(string), incrementValue, client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update anime list rating or anime list not found"})
		return
	}
	fmt.Println(result)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "List rating updated successfully"})
}

func GetAnimeListsByAnimeId(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	animeList, err := database.GetAllAnimeListsByAnimeId(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterListsByCharacterId(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	animeList, err := database.GetAllCharacterListsByCharacterId(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetAnimeListsByUserId(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	animeList, err := database.GetAllAnimeListsByUserId(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

func GetCharacterListsByUserId(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	animeList, err := database.GetAllCharacterListsByUserId(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, animeList)
}

type AnimeListsWithContent struct {
	AnimeList database.AnimeList `json:"anime_list"`
	Anime     []database.Anime   `json:"anime"`
}

func GetAnimeListsWithAnime(c *gin.Context, client *mongo.Client) {
	animeList, err := database.GetAllAnimeLists(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime lists"})
		return
	}
	var result []AnimeListsWithContent
	for _, list := range animeList {
		anime, err := database.GetAnimeFromListToDisplay(list, client)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime lists"})
			return
		}
		result = append(result, AnimeListsWithContent{AnimeList: list, Anime: anime})
	}
	c.IndentedJSON(http.StatusOK, result)
}

type CharactersListsWithContent struct {
	CharactersList database.CharacterList `json:"characters_list"`
	Characters     []database.Character   `json:"characters"`
}

func GetCharacterListsWithCharacters(c *gin.Context, client *mongo.Client) {
	characterList, err := database.GetAllCharacterLists(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve character lists"})
		return
	}
	var result []CharactersListsWithContent
	for _, list := range characterList {
		characters, err := database.GetCharactersFromListToDisplay(list, client)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve character lists"})
			return
		}
		result = append(result, CharactersListsWithContent{CharactersList: list, Characters: characters})
	}
	c.IndentedJSON(http.StatusOK, result)
}
