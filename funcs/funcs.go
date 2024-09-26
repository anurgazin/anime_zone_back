package funcs

import (
	"fmt"
	"os"

	// "io"
	"net/http"

	"github.com/joho/godotenv"
)

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
