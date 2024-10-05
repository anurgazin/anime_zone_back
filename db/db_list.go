package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAnimeList(listName string, id string) (interface{}, error) {
	var animeList AnimeList
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("AnimeList")
	animeList.ID = primitive.NewObjectID()
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UserId format: %w", err)
	}
	animeList.UserID = userId
	animeList.Name = listName
	insertResult, err := collection.InsertOne(context.TODO(), animeList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func CreateCharacterList(listName string, id string) (interface{}, error) {
	var characterList CharacterList
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("CharacterList")
	characterList.ID = primitive.NewObjectID()
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UserId format: %w", err)
	}
	characterList.UserID = userId
	characterList.Name = listName
	insertResult, err := collection.InsertOne(context.TODO(), characterList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}
