package funcs

import (
	database "anime_zone/back_end/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

var anime = database.SampleAnime

func GetAnime(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, anime)
}

func GetAnimeById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range anime {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anime not found"})
}

func GetAnimeByTitle(c *gin.Context) {
	title := c.Param("title")

	for _, a := range anime {
		if a.Title == title {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "anime not found"})
}

func PostAnime(c *gin.Context) {
	var newAnime database.Anime

	if err := c.BindJSON(&newAnime); err != nil {
		return
	}

	if checkAnimeExistsById(newAnime.ID) {
		// If any anime exists, return an error message
		c.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime already exists in our db: " + newAnime.Title})
		return
	}

	anime = append(anime, newAnime)
	c.IndentedJSON(http.StatusCreated, newAnime)
}

func PutAnime(c *gin.Context) {
	var updatedAnime database.Anime

	id := c.Param("id")

	if err := c.BindJSON(&updatedAnime); err != nil {
		return
	}

	for i, a := range anime {
		if a.ID == id {
			anime[i].Title = updatedAnime.Title
			anime[i].Description = updatedAnime.Description
			anime[i].Duration = updatedAnime.Duration
			anime[i].ESRB = updatedAnime.ESRB
			anime[i].Episodes = updatedAnime.Episodes
			anime[i].Rating = updatedAnime.Rating
			anime[i].ReleaseDate = updatedAnime.ReleaseDate
			anime[i].Status = updatedAnime.Status
			anime[i].Studio = updatedAnime.Studio
			anime[i].Type = updatedAnime.Type
			anime[i].Genre = updatedAnime.Genre
			c.JSON(http.StatusOK, gin.H{"message": "Anime updated: " + a.ID})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Such Anime not found: " + updatedAnime.Title})
}
