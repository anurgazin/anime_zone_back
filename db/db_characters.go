package database

import (
	azureblob "anime_zone/back_end/azure_blob"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UploadCharacter(character Character) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Characters")
	character.ID = primitive.NewObjectID()
	insertResult, err := collection.InsertOne(context.TODO(), character)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func GetAllCharacters() ([]Character, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Characters")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var result []Character
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all characters")
	return result, nil
}

func GetCharacterById(id string) (*Character, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Characters")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var result Character
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no anime found with the given ID")
		}
		return nil, fmt.Errorf("error finding anime: %w", err)
	}

	// Return the found anime
	return &result, nil
}

func UpdateCharacter(id string, updatedCharacter Character) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Characters")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	// Create the update document (bson.M) with the fields you want to update.
	// $set operator is used to update the provided fields
	update := bson.M{
		"$set": bson.M{
			"first_name": updatedCharacter.FirstName,
			"last_name":  updatedCharacter.LastName,
			"age":        updatedCharacter.Age,
			"bio":        updatedCharacter.Bio,
			"from_anime": updatedCharacter.FromAnime,
			"gender":     updatedCharacter.Gender,
			"logo":       updatedCharacter.Logo,
			"media":      updatedCharacter.Media,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("could not update character: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no character found with the given ID")
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.ModifiedCount)
	return result, nil
}

func DeleteCharacter(id string) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Characters")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	character, err := GetCharacterById(id)

	if err != nil {
		return nil, fmt.Errorf("no character found with the given ID")
	}

	deleteLogoResult, err := azureblob.DeleteFile(character.Logo)
	if err != nil {
		return nil, err
	}
	fmt.Println(deleteLogoResult)
	for i := range character.Media {
		deleteMediaResult, err := azureblob.DeleteFile(character.Media[i])
		if err != nil {
			return nil, err
		}
		fmt.Println(deleteMediaResult)
	}

	deleteCommentResult, err := DeleteCommentByContentId(id, "character")
	if err != nil {
		return nil, fmt.Errorf("error during deleting comments")
	}
	fmt.Println(deleteCommentResult)

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete character: %w", err)
	}

	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("no character found with the given ID")
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.DeletedCount)
	return result, nil
}
