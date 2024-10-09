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
	"strconv"
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

func handleImageUploader(file *multipart.FileHeader, title, suffix string) (string, error) {
	fileParts := strings.Split(file.Filename, ".")
	extension := fileParts[len(fileParts)-1]
	fileName := strings.ReplaceAll(title, " ", "_") + suffix + "." + extension

	byteContainer, err := ConvertFileHeaderToBytes(file)
	if err != nil {
		log.Println("ConvertFile error:")
		log.Println(err.Error())
		return "", err
	}

	fileUrl, err := azureblob.UploadFile(fileName, byteContainer)
	if err != nil {
		log.Println("UploadFile error:")
		log.Println(err.Error())
		return "", err
	}
	return fileUrl, nil
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

	logoUrl, err := handleImageUploader(newAnimeUploader.Logo, newAnimeUploader.Title, "_Logo")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	mediaURLs := []string{newAnimeUploader.Link}
	for i := range newAnimeUploader.Media {
		url, err := handleImageUploader(newAnimeUploader.Media[i], newAnimeUploader.Title, "_Media_"+strconv.Itoa(i+1))
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
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, anime)
}
