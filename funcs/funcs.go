package funcs

import (
	"fmt"
	// "io"
	"net/http"
)

// Function to check if an anime exists by making a GET request to /anime/:title
func checkAnimeExistsByTitle(title string) bool {
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
func checkAnimeExistsById(id string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/anime/id/" + id
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		return false
	}

	return true
}

// Function to check if an anime exists by making a GET request to /anime/:title
func checkCharactersExistsById(id string) bool {
	// Make the GET request to the /anime/:title endpoint
	url := "http://localhost:8080/characters/" + id
	resp, err := http.Get(url)

	// If there is an error or the status code is not 200 (OK), return false
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
