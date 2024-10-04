package middleware

import (
	"anime_zone/back_end/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthToken(c *gin.Context) {
	tokenString := c.GetHeader("token")
	if tokenString == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Token is missing"})
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Verify the token
	token, err := jwt.VerifyToken(tokenString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)

	c.Next()
}
