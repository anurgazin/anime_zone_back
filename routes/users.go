package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"anime_zone/back_end/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	var newUser database.User
	var newUserUploader database.UserUploader

	if err := c.ShouldBind(&newUserUploader); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	logoUrl, err := funcs.HandleImageUploader(newUserUploader.Logo, newUserUploader.Username, "_Logo")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	newUser = database.User{
		Email:    newUserUploader.Email,
		Username: newUserUploader.Username,
		Password: newUserUploader.Password,
		Role:     newUserUploader.Role,
		Bio:      newUserUploader.Bio,
		Logo:     logoUrl,
	}

	insertedID, err := database.RegisterUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Registered next user: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	result, err := database.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	userToken, err := jwt.CreateToken(result.(database.User))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"token": userToken})
}

func PutUser(c *gin.Context) {
	var updatedUser database.User

	id := c.Param("id")

	if err := c.BindJSON(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	result, err := database.EditUser(id, updatedUser)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}
