package routes

import (
	database "anime_zone/back_end/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	var newUser database.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	insertedID, err := database.RegisterUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Registered next user: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}
