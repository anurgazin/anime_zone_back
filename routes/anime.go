package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"strconv"
)

func GetAnime(c *gin.Context, client *mongo.Client) {
	anime, err := database.GetAllAnime(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime"})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetAnimeById(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	anime, err := database.GetAnimeById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetAnimeByTitle(c *gin.Context, client *mongo.Client) {
	title := c.Param("title")

	anime, err := database.GetAnimeByTitle(title, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func PostAnime(c *gin.Context, client *mongo.Client) {
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

	insertedID, err := database.UploadAnime(newAnime, client)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Added next anime: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

func PutAnime(c *gin.Context, client *mongo.Client) {
	var updatedAnime database.Anime

	id := c.Param("id")

	if err := c.BindJSON(&updatedAnime); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid anime data"})
		return
	}

	result, err := database.UpdateAnime(id, updatedAnime, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}

func DeleteAnime(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	anime, err := database.DeleteAnime(id, client)

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

func RateAnime(c *gin.Context, client *mongo.Client) {
	// Parse the rating request body
	anime_id := c.Param("id")
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
	var ratingRequest RatingRequest
	if err := c.ShouldBindJSON(&ratingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.PostRating(anime_id, user_id.(string), username.(string), ratingRequest.Score, ratingRequest.Review, client)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit rating"})
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"message": "Rating submitted successfully!"})
}

func UpdateRating(c *gin.Context, client *mongo.Client) {
	// Parse the rating request body
	review_id := c.Param("id")
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

	result, err := database.EditRating(review_id, user_id.(string), ratingRequest.Score, ratingRequest.Review, client)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rating"})
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"message": "Rating updated successfully!"})
}

func GetAnimeRatingById(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	anime, err := database.GetAnimeRatingById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetAnimeRatingByUser(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	anime, err := database.GetAnimeRatingByUserId(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetHighestRatedAnime(c *gin.Context, client *mongo.Client) {
	anime, err := database.GetHighestRatedAnime(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime"})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetMostPopularAnime(c *gin.Context, client *mongo.Client) {
	anime, err := database.GetMostPopularAnime(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve anime"})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}

func GetSimilarAnime(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	anime, err := database.GetAnimeById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	result, err := database.GetSimilarAnime(anime.Genre, anime.Studio, id, client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

type AnimeDetailsResult struct {
	Anime        database.Anime       `json:"anime"`
	SimilarAnime []database.Anime     `json:"similar_anime"`
	AnimeReviews []database.Rating    `json:"reviews"`
	Characters   []database.Character `json:"characters"`
	AnimeList    []database.AnimeList `json:"anime_list"`
	Comments     []database.Comment   `json:"comments"`
}

func GetAnimeDetails(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	anime, err := database.GetAnimeById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	similarAnime, err := database.GetSimilarAnime(anime.Genre, anime.Studio, id, client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	animeReviews, err := database.GetAnimeRatingById(id, client)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	characters, err := database.GetAllCharactersFromAnime(id, client)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	animeList, err := database.GetAllAnimeListsByAnimeId(id, client)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	comment, err := database.GetAllCommentsForContent("anime", id, client)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	var result AnimeDetailsResult
	result.Anime = *anime
	result.SimilarAnime = similarAnime
	result.AnimeReviews = animeReviews
	result.Characters = characters
	result.AnimeList = animeList
	result.Comments = comment

	c.IndentedJSON(http.StatusOK, result)
}

type AnimeListDetails struct {
	Anime     []database.Anime   `json:"anime"`
	AnimeList database.AnimeList `json:"anime_list"`
	Comments  []database.Comment `json:"comments"`
}

func GetAllAnimeFromList(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	list, err := database.GetAnimeListById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	anime, err := database.GetAllAnimeFromList(*list, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	comment, err := database.GetAllCommentsForContent("anime_list", id, client)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	var result AnimeListDetails
	result.Anime = anime
	result.AnimeList = *list
	result.Comments = comment

	c.IndentedJSON(http.StatusOK, result)
}

func GetAnimeReviewsById(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	review, err := database.GetReviewById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, review)
}
