package jwt

import (
	"fmt"
	"time"

	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"

	"github.com/golang-jwt/jwt/v5"
)

// Function to create JWT tokens with claims
func CreateToken(user database.User) (string, error) {
	// Create a new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username, // Subject (user identifier)
		"id":       user.ID,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":      time.Now().Unix(),                // Issued at
	})

	// Print information about the created token
	fmt.Printf("Token claims added: %+v\n", token)
	tokenString, err := token.SignedString(funcs.GoDotEnvVariable("JWT_SECRET"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return funcs.GoDotEnvVariable("JWT_SECRET"), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
