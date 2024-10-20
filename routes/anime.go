package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAnime(c *gin.Context) {
	anime, err := database.GetAllAnime()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime"})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetAnimeById(c *gin.Context) {
	id := c.Param("id")

	anime, err := database.GetAnimeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetAnimeByTitle(c *gin.Context) {
	title := c.Param("title")

	anime, err := database.GetAnimeByTitle(title)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func PostAnime(c *gin.Context) {
	var newAnime database.Anime
	var newAnimeUploader database.AnimeUploader

	if err := c.ShouldBind(&newAnimeUploader); err != nil {
		log.Println("Binding error:")
		log.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid anime data"})
		return
	}

	logoUrl, err := funcs.HandleImageUploader(newAnimeUploader.Logo, newAnimeUploader.Title, "_Logo")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	mediaURLs := []string{newAnimeUploader.Link}
	for i := range newAnimeUploader.Media {
		url, err := funcs.HandleImageUploader(newAnimeUploader.Media[i], newAnimeUploader.Title, "_Media_"+strconv.Itoa(i+1))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		mediaURLs = append(mediaURLs, url)
	}

	newAnime.Title = newAnimeUploader.Title
	newAnime.ReleaseDate = newAnimeUploader.ReleaseDate
	newAnime.Rating = newAnimeUploader.Rating
	newAnime.Genre = newAnimeUploader.Genre
	newAnime.Type = newAnimeUploader.Type
	newAnime.Episodes = newAnimeUploader.Episodes
	newAnime.Description = newAnimeUploader.Description
	newAnime.Studio = newAnimeUploader.Studio
	newAnime.Duration = newAnimeUploader.Duration
	newAnime.Status = newAnimeUploader.Status
	newAnime.ESRB = newAnimeUploader.ESRB
	newAnime.Logo = logoUrl
	newAnime.Media = mediaURLs

	insertedID, err := database.UploadAnime(newAnime)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Added next anime: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

func PutAnime(c *gin.Context) {
	var updatedAnime database.Anime

	id := c.Param("id")

	if err := c.BindJSON(&updatedAnime); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid anime data"})
		return
	}

	result, err := database.UpdateAnime(id, updatedAnime)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}

func DeleteAnime(c *gin.Context) {
	id := c.Param("id")
	anime, err := database.DeleteAnime(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

// type UpdateRatingAction struct {
// 	Action string `json:"action" form:"action"` // "increment" or "decrement"
// }

// func UpdateAnimeRating(c *gin.Context) {
// 	id := c.Param("id")
// 	var updateData UpdateRatingAction

// 	// Bind the action (increment or decrement) from the request
// 	if err := c.ShouldBindJSON(&updateData); err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
// 		return
// 	}
// 	// Determine the increment value based on the action
// 	var incrementValue float64
// 	if updateData.Action == "increment" {
// 		incrementValue = 1.0
// 	} else if updateData.Action == "decrement" {
// 		incrementValue = -1.0
// 	} else {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid action. Use 'increment' or 'decrement'."})
// 		return
// 	}

// 	result, err := database.UpdateAnimeRating(id, incrementValue)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update anime rating or anime not found"})
// 		return
// 	}
// 	fmt.Println(result)
// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Anime rating updated successfully"})
// }
