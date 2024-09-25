package funcs

import (
	"anime_zone/back_end/database"

	// "fmt"
	// "io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var characters = database.SampleCharacters

func GetCharacters(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, characters)
}

func GetCharactersById(g *gin.Context) {
	id := g.Param("id")

	for _, c := range characters {
		if c.ID == id {
			g.IndentedJSON(http.StatusOK, c)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "character not found"})
}

func PostCharacters(g *gin.Context) {
	var newCharacter database.Character

	// Bind the incoming JSON to the newCharacter struct
	if err := g.BindJSON(&newCharacter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character data"})
		return
	}

	if checkCharactersExistsById(newCharacter.ID) {
		// If such character exists, return an error message
		g.JSON(http.StatusBadRequest, gin.H{"error": "Such Character already exists in our db: " + newCharacter.FirstName + " " + newCharacter.LastName})
		return
	}

	// Check if all anime in FromAnime exist
	for _, animeTitle := range newCharacter.FromAnime {
		if !checkAnimeExistsByTitle(animeTitle) {
			// If any anime doesn't exist, return an error message
			g.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime doesn't exist in our db: " + animeTitle})
			return
		}
	}

	// If all anime exist, add the new character
	characters = append(characters, newCharacter)

	// Respond with the created character
	g.IndentedJSON(http.StatusCreated, newCharacter)
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
	for _, animeTitle := range updatedCharacter.FromAnime {
		if !checkAnimeExistsByTitle(animeTitle) {
			// If any anime doesn't exist, return an error message
			g.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime doesn't exist in our db: " + animeTitle})
			return
		}
	}

	for i, c := range characters {
		if c.ID == id {
			characters[i].FirstName = updatedCharacter.FirstName
			characters[i].LastName = updatedCharacter.LastName
			characters[i].Bio = updatedCharacter.Bio
			characters[i].Age = updatedCharacter.Age
			characters[i].Gender = updatedCharacter.Gender
			characters[i].FromAnime = updatedCharacter.FromAnime
			characters[i].Status = updatedCharacter.Status

			g.JSON(http.StatusOK, gin.H{"message": "Character updated: " + c.ID})
			return
		}
	}

	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Such Character not found: " + updatedCharacter.FirstName + " " + updatedCharacter.LastName})
}
