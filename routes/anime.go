package routes

import (
	azureblob "anime_zone/back_end/azure_blob"
	database "anime_zone/back_end/db"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ConvertFileHeaderToBytes takes a *multipart.FileHeader and converts it into a byte array
func ConvertFileHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a buffer to hold the file's contents
	var buf bytes.Buffer
	// Copy the file's contents into the buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return nil, err
	}

	// Return the byte array
	return buf.Bytes(), nil
}

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

	if err := c.Bind(&newAnimeUploader); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid anime data"})
		return
	}

	log.Println(newAnimeUploader)
	filename := strings.Split(newAnimeUploader.Logo.Filename, ".")
	log.Println(filename[len(filename)-1])
	logoFileName := strings.ReplaceAll(newAnimeUploader.Title, " ", "_") + "_Logo." + filename[len(filename)-1]

	byteContainer, err := ConvertFileHeaderToBytes(newAnimeUploader.Logo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert file"})
		return
	}

	logoUrl, err := azureblob.UploadFile(logoFileName, byteContainer)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(logoUrl)

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
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}
