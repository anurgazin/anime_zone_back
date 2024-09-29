package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"fmt"

	// "fmt"
	// "io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCharacters(g *gin.Context) {
	characters, err := database.GetAllCharacters()
	if err != nil {
		g.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve characters"})
		return
	}
	g.IndentedJSON(http.StatusOK, characters)
}

func GetCharactersById(g *gin.Context) {
	id := g.Param("id")
	character, err := database.GetCharacterById(id)

	if err != nil {
		g.IndentedJSON(http.StatusNotFound, gin.H{"message": "character not found"})
		return
	}
	g.IndentedJSON(http.StatusOK, character)
}

func PostCharacters(g *gin.Context) {
	var newCharacter database.Character

	// Bind the incoming JSON to the newCharacter struct
	if err := g.BindJSON(&newCharacter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character data"})
		return
	}

	if funcs.CheckCharactersExistsById(primitive.ObjectID(newCharacter.ID).String()) {
		// If such character exists, return an error message
		g.JSON(http.StatusBadRequest, gin.H{"error": "Such Character already exists in our db: " + newCharacter.FirstName + " " + newCharacter.LastName})
		return
	}

	// Check if all anime in FromAnime exist
	for _, animeId := range newCharacter.FromAnime {
		if !funcs.CheckAnimeExistsById(animeId.Hex()) {
			// If any anime doesn't exist, return an error message
			g.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime doesn't exist in our db: " + animeId.Hex()})
			return
		}
	}

	insertedID, err := database.UploadCharacter(newCharacter)
	if err != nil {
		g.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Added next character: %v", insertedID)
	g.IndentedJSON(http.StatusCreated, result)
}

func PutCharacters(g *gin.Context) {
	var updatedCharacter database.Character
	id := g.Param("id")

	// Bind the incoming JSON to the updatedCharacter struct
	if err := g.BindJSON(&updatedCharacter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character data"})
		return
	}

	// Check if all anime in FromAnime exist
	for _, animeId := range updatedCharacter.FromAnime {
		if !funcs.CheckAnimeExistsById(animeId.Hex()) {
			// If any anime doesn't exist, return an error message
			g.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime doesn't exist in our db: " + animeId.Hex()})
			return
		}
	}

	result, err := database.UpdateCharacter(id, updatedCharacter)

	if err != nil {
		g.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	g.IndentedJSON(http.StatusNotFound, gin.H{"message": result})
}

func DeleteCharacter(g *gin.Context) {
	id := g.Param("id")
	character, err := database.DeleteCharacter(id)

	if err != nil {
		g.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	g.IndentedJSON(http.StatusOK, character)
}
