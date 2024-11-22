package middleware

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_lib "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthToken(c *gin.Context) {
	tokenString := c.GetHeader("Auth")
	if tokenString == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Token is missing"})
		c.Abort()
		return
	}

	token, err := jwt.VerifyToken(tokenString, false)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err})
		c.Abort()
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	fmt.Println()

	c.Set("token", token)
	c.Set("claims", token.Claims)
	c.Set("id", token.Claims.(jwt_lib.MapClaims)["id"].(string))
	c.Set("role", token.Claims.(jwt_lib.MapClaims)["role"])
	c.Set("username", token.Claims.(jwt_lib.MapClaims)["username"])
	c.Next()
}

func RefreshToken(c *gin.Context, client *mongo.Client) {
	refreshTokenString := c.GetHeader("RefreshToken")
	if refreshTokenString == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Refresh token is missing"})
		c.Abort()
		return
	}

	token, err := jwt.VerifyToken(refreshTokenString, true) // Verify as refresh token
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt_lib.MapClaims)
	if !ok || !token.Valid {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	// Generate new access and refresh tokens
	id := claims["id"].(string)
	user, err := database.GetUser(id, client) // Replace with your function to fetch user data
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		c.Abort()
		return
	}
	userTokens, err := jwt.CreateTokens(*user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tokens"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  userTokens.AccessToken,
		"refresh_token": userTokens.RefreshToken,
	})
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
