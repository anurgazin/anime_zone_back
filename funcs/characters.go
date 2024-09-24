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

// Function to check if an anime exists by making a GET request to /anime/:title
func checkAnimeExists(title string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/anime/title/" + title
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	// Optionally read the response body to ensure the anime exists
	// defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println("Response from anime endpoint:", string(body))

	// If everything is fine, return true
	return true
}

func PostCharacters(g *gin.Context) {
	var newCharacter database.Character

	// Bind the incoming JSON to the newCharacter struct
	if err := g.BindJSON(&newCharacter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character data"})
		return
	}

	// Check if all anime in FromAnime exist
	for _, animeTitle := range newCharacter.FromAnime {
		if !checkAnimeExists(animeTitle) {
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
