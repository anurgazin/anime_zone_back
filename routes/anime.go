package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
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
	newAnime.AverageRating = newAnimeUploader.AverageRating
	newAnime.RatingCount = newAnimeUploader.RatingCount
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

type RatingRequest struct {
	Score  float64 `json:"score" binding:"required"`
	Review string  `json:"review"` // optional
}

func RateAnime(c *gin.Context) {
	// Parse the rating request body
	anime_id := c.Param("id")
	user_id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	var ratingRequest RatingRequest
	if err := c.ShouldBindJSON(&ratingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.PostRating(anime_id, user_id.(string), ratingRequest.Score, ratingRequest.Review)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit rating"})
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"message": "Rating submitted successfully!"})
}
