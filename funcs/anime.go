package funcs

import (
	"anime_zone/back_end/database"
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

func PostAnime(c *gin.Context) {
	var newAnime database.Anime

	if err := c.BindJSON(&newAnime); err != nil {
		return
	}

	anime = append(anime, newAnime)
	c.IndentedJSON(http.StatusCreated, newAnime)
}
