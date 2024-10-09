package routes

import (
	azureblob "anime_zone/back_end/azure_blob"
	database "anime_zone/back_end/db"
	"fmt"
	"net/http"

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

	if err := c.BindJSON(&newAnime); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid anime data"})
		return
	}

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

// CheckImage handles file upload via multipart form data
func CheckImage(c *gin.Context) {
	// Retrieve the file from the form-data
	file, header, err := c.Request.FormFile("new_image")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image: " + err.Error()})
		return
	}
	defer file.Close()

	// Get the file name from the header
	fileName := header.Filename

	// Read the file content
	fileBytes := make([]byte, header.Size)
	_, err = file.Read(fileBytes)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file: " + err.Error()})
		return
	}

	// Upload the file to Azure Blob Storage
	fileURL, err := azureblob.UploadFile(fileName, fileBytes)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the URL of the uploaded file
	c.IndentedJSON(http.StatusOK, gin.H{"file_url": fileURL})
}
