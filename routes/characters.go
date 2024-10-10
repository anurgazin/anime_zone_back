package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"fmt"
	"log"
	"strconv"

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
	var newCharacterUploader database.CharacterUploader

	// Bind the incoming JSON to the newCharacterUploader struct
	if err := g.ShouldBind(&newCharacterUploader); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character data"})
		return
	}

	// Check if all anime in FromAnime exist
	var fromAnime = []primitive.ObjectID{}
	for _, animeId := range newCharacterUploader.FromAnime {
		if !funcs.CheckAnimeExistsById(animeId) {
			// If any anime doesn't exist, return an error message
			g.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime doesn't exist in our db: " + animeId})
			return
		}
		id, err := primitive.ObjectIDFromHex(animeId)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid ObjectID format " + animeId})
		}
		fromAnime = append(fromAnime, id)
	}

	var title string = newCharacterUploader.FirstName + " " + newCharacterUploader.LastName

	logoUrl, err := funcs.HandleImageUploader(newCharacterUploader.Logo, title, "_Logo")
	if err != nil {
		g.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	mediaURLs := []string{}
	for i := range newCharacterUploader.Media {
		url, err := funcs.HandleImageUploader(newCharacterUploader.Media[i], title, "_Media_"+strconv.Itoa(i+1))
		if err != nil {
			g.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		mediaURLs = append(mediaURLs, url)
	}

	newCharacter = database.Character{
		FirstName: newCharacterUploader.FirstName,
		LastName:  newCharacterUploader.LastName,
		Age:       newCharacterUploader.Age,
		FromAnime: fromAnime,
		Gender:    newCharacterUploader.Gender,
		Bio:       newCharacterUploader.Bio,
		Status:    newCharacterUploader.Status,
		Logo:      logoUrl,
		Media:     mediaURLs,
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
