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

func RegisterUser(user User, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Users")
	pwd := user.Password
	password, err := HashPassword(pwd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user.Password = password
	user.ID = primitive.NewObjectID()

	if user.Role != "admin" {
		user.Role = "guest"
	}

	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func LoginUser(email string, password string, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Users")
	filter := bson.M{"email": email}

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

func EditUser(id string, updatedUser User, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}
	pwd := updatedUser.Password
	password, err := HashPassword(pwd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	updatedUser.Password = password

	update := bson.M{
		"$set": bson.M{
			"username": updatedUser.Username,
			"bio":      updatedUser.Bio,
			"logo":     updatedUser.Logo,
			"password": updatedUser.Password,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("could not edit user: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no users found with the given ID")
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.ModifiedCount)
	return result, nil
}

func GetUser(id string, client *mongo.Client) (*User, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var result User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no user found with the given ID")
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	// Return the found user
	return &result, nil
}
