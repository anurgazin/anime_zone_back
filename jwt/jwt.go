package jwt

import (
	"fmt"
	"time"

	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(funcs.GoDotEnvVariable("JWT_SECRET"))
var refreshSecret = []byte(funcs.GoDotEnvVariable("JWT_REFRESH_SECRET"))

// Function to create access and refresh tokens
func CreateTokens(user database.User) (*database.TokenPair, error) {
	// Create Access Token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"role":     user.Role,
		"exp":      time.Now().Add(15 * time.Minute).Unix(), // Expiration time for access token
		"iat":      time.Now().Unix(),                       // Issued at
	})

	accessTokenString, err := accessToken.SignedString(secret)
	if err != nil {
		return nil, err
	}

	// Create Refresh Token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(), // Longer expiration time for refresh token
		"iat": time.Now().Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return nil, err
	}

	tokenPair := &database.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return tokenPair, nil
}

// Function to verify tokens
func VerifyToken(tokenString string, isRefresh bool) (*jwt.Token, error) {
	var key []byte
	if isRefresh {
		key = refreshSecret
	} else {
		key = secret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
