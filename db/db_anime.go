package database

import (
	azureblob "anime_zone/back_end/azure_blob"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UploadAnime(anime Anime) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Anime")
	anime.ID = primitive.NewObjectID()
	if anime.Media == nil {
		anime.Media = []string{}
	}
	insertResult, err := collection.InsertOne(context.TODO(), anime)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func GetAllAnime() ([]Anime, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Anime")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var result []Anime
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all anime")
	return result, nil
}

func GetAnimeById(id string) (*Anime, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Anime")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var result Anime
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

func GetAnimeByTitle(title string) (*Anime, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Anime")

	filter := bson.M{"title": title}

	var result Anime
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no anime found with the given title")
		}
		return nil, fmt.Errorf("error finding anime: %w", err)
	}

	// Return the found anime
	return &result, nil
}

func UpdateAnime(id string, updatedAnime Anime) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Anime")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	// Create the update document (bson.M) with the fields you want to update.
	// $set operator is used to update the provided fields
	update := bson.M{
		"$set": bson.M{
			"title":        updatedAnime.Title,
			"release_date": updatedAnime.ReleaseDate,
			"rating":       updatedAnime.Rating,
			"genre":        updatedAnime.Genre,
			"type":         updatedAnime.Type,
			"episodes":     updatedAnime.Episodes,
			"description":  updatedAnime.Description,
			"studio":       updatedAnime.Studio,
			"duration":     updatedAnime.Duration,
			"status":       updatedAnime.Status,
			"esrb":         updatedAnime.ESRB,
			"logo":         updatedAnime.Logo,
			"media":        updatedAnime.Media,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("could not update anime: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no anime found with the given ID")
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.ModifiedCount)
	return result, nil
}

func DeleteAnime(id string) (interface{}, error) {
	client := RunMongo()
	anime_collection := client.Database("Anime-Zone").Collection("Anime")
	character_collection := client.Database("Anime-Zone").Collection("Characters")

	anime_get, err := GetAnimeById(id)
	if err != nil {
		return nil, fmt.Errorf("no anime found with the given ID")
	}

	deleteLogoResult, err := azureblob.DeleteFile(anime_get.Logo)
	if err != nil {
		return nil, err
	}
	fmt.Println(deleteLogoResult)
	media := anime_get.Media[1:]
	for i := range media {
		deleteMediaResult, err := azureblob.DeleteFile(media[i])
		if err != nil {
			return nil, err
		}
		fmt.Println(deleteMediaResult)
	}

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}
	filter := bson.M{"_id": objID}
	anime_result, err := anime_collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete anime: %w", err)
	}

	if anime_result.DeletedCount == 0 {
		return nil, fmt.Errorf("no anime found with the given ID")
	}

	// Delete characters where `from_anime` contains the anime ID
	character_filter := bson.M{"from_anime": bson.M{"$in": []primitive.ObjectID{objID}}}
	character_cursor, err := character_collection.Find(context.TODO(), character_filter)
	if err != nil {
		return nil, fmt.Errorf("could not get characters: %w", err)
	}
	var characters_get_result []Character
	if err := character_cursor.All(context.TODO(), &characters_get_result); err != nil {
		return nil, err
	}
	var delete_character_result interface{}
	for i := range characters_get_result {
		delete_character_result, err = DeleteCharacter(characters_get_result[i].ID.Hex())
		if err != nil {
			return nil, err
		}
	}
	comment_result, err := DeleteCommentByContentId(id, "anime")
	if err != nil {
		return nil, fmt.Errorf("error during deleting comments")
	}

	return map[string]interface{}{
		"deleted_anime_count":     anime_result.DeletedCount,
		"deleted_character_count": delete_character_result,
		"deleted_comment_count":   comment_result,
	}, nil
}

// func UpdateAnimeRating(id string, value float64) (interface{}, error) {
// 	client := RunMongo()
// 	collection := client.Database("Anime-Zone").Collection("Anime")

// 	// Convert the string ID to ObjectID
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
// 	}

// 	filter := bson.M{"_id": objID}
// 	update := bson.M{"$inc": bson.M{"rating": value}}
// 	result, err := collection.UpdateOne(context.TODO(), filter, update)

// 	if err != nil {
// 		return nil, fmt.Errorf("could not update anime rating: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return nil, fmt.Errorf("no anime found with the given ID")
// 	}

// 	fmt.Printf("Successfully updated %v document(s)\n", result.ModifiedCount)
// 	return result, nil
// }
