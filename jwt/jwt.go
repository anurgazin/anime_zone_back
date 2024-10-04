package jwt

import (
	"fmt"
	"time"

	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(funcs.GoDotEnvVariable("JWT_SECRET"))

// Function to create JWT tokens with claims
func CreateToken(user database.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":      time.Now().Unix(),                // Issued at
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Function to verify JWT tokens
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
