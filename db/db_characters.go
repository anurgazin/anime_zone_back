package database

import (
	azureblob "anime_zone/back_end/azure_blob"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UploadCharacter(character Character, client *mongo.Client) (interface{}, error) {
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

func GetAllCharacters(client *mongo.Client) ([]Character, error) {
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

func GetCharacterById(id string, client *mongo.Client) (*Character, error) {
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

func UpdateCharacter(id string, updatedCharacter Character, client *mongo.Client) (interface{}, error) {
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

func DeleteCharacter(id string, client *mongo.Client) (interface{}, error) {
	collection := client.Database("Anime-Zone").Collection("Characters")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	character, err := GetCharacterById(id, client)

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

	deleteCommentResult, err := DeleteCommentByContentId(id, "character", client)
	if err != nil {
		return nil, fmt.Errorf("error during deleting comments")
	}
	fmt.Println(deleteCommentResult)

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete character: %w", err)
	}
	return result, nil
}

func GetAllCharactersFromAnime(anime_id string, client *mongo.Client) ([]Character, error) {
	collection := client.Database("Anime-Zone").Collection("Characters")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(anime_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"from_anime.id": objID}

	cursor, err := collection.Find(context.TODO(), filter)
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

func GetCharactersFirstName(client *mongo.Client) ([]Character, error) {
	collection := client.Database("Anime-Zone").Collection("Characters")

	opts := options.Find().SetSort(bson.D{{Key: "last_name", Value: 1}}).SetLimit(10)

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
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

func GetAllCharactersFromList(characters_list CharacterList, client *mongo.Client) ([]Character, error) {
	anime_collection := client.Database("Anime-Zone").Collection("Characters")

	filter := bson.M{"_id": bson.M{"$in": characters_list.CharacterList}}

	cursor, err := anime_collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find character from list: %w", err)
	}

	var result []Character
	if err := cursor.All(context.TODO(), &result); err != nil {
		return nil, fmt.Errorf("failed to decode characters from list: %w", err)
	}
	fmt.Println("Retrieved all characters from list")
	return result, nil
}
