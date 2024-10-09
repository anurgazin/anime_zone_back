package middleware

import (
	"anime_zone/back_end/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_lib "github.com/golang-jwt/jwt/v5"
)

func AuthToken(c *gin.Context) {
	tokenString := c.GetHeader("Auth")
	if tokenString == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Token is missing"})
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	token, err := jwt.VerifyToken(tokenString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	fmt.Println()

	c.Set("token", token)
	c.Set("claims", token.Claims)
	c.Set("id", token.Claims.(jwt_lib.MapClaims)["id"].(string))
	c.Next()
}

func IsAdmin(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		c.Abort()
		return
	}

	role, ok := claims.(jwt_lib.MapClaims)["role"].(string)

	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}
	if role != "admin" {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		c.Abort()
		return
	}
	c.Next()
}
