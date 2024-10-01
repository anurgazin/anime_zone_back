package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(user User) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Users")
	pwd := user.Password
	password, err := HashPassword(pwd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user.Password = password
	user.ID = primitive.NewObjectID()

	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func LoginUser(username string, password string) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Users")
	filter := bson.M{"username": username}

	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no user found with the given username")
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	match := VerifyPassword(password, result.Password)
	if !match {
		return nil, fmt.Errorf("incorrect Password")
	}

	return result, nil
}
