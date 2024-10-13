package funcs

import (
	azureblob "anime_zone/back_end/azure_blob"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
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

func HandleImageUploader(file *multipart.FileHeader, title, suffix string) (string, error) {
	var contentType string
	fileParts := strings.Split(file.Filename, ".")
	extension := fileParts[len(fileParts)-1]
	fileName := strings.ReplaceAll(title, " ", "_") + suffix + "." + extension

	byteContainer, err := ConvertFileHeaderToBytes(file)
	if err != nil {
		log.Println("ConvertFile error:")
		log.Println(err.Error())
		return "", err
	}
	if extension == "jpg" {
		contentType = "image/jpeg"
	}
	if extension == "png" {
		contentType = "image/png"
	}
	if extension == "webp" {
		contentType = "image/webp"
	}
	fileUrl, err := azureblob.UploadFile(fileName, byteContainer, contentType)
	if err != nil {
		log.Println("UploadFile error:")
		log.Println(err.Error())
		return "", err
	}
	return fileUrl, nil
}

// Function to get ENV variables
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

// Function to check if an anime exists by making a GET request to /anime/:title
func CheckAnimeExistsByTitle(title string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/anime/title/" + title
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

// Function to check if an anime exists by making a GET request to /anime/:id
func CheckAnimeExistsById(id string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/anime/id/" + id
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

// Function to check if an anime exists by making a GET request to /anime/:title
func CheckCharactersExistsById(id string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/characters/" + id
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

// Function to check if a character list exists by making a GET request to /list/characters/:id
func CheckCharacterListExistsById(id string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/list/characters/" + id
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

// Function to check if an anime list exists by making a GET request to /list/anime/:id
func CheckAnimeListExistsById(id string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/list/anime/" + id
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
