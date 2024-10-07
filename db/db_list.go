package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	animeList.AnimeList = []primitive.ObjectID{}
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
	characterList.CharacterList = []primitive.ObjectID{}
	insertResult, err := collection.InsertOne(context.TODO(), characterList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func AddAnimeToList(listId string, userId string, animeId string) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	uID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (UserId): %w", err)
	}

	lID, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (ListId): %w", err)
	}

	filter := bson.M{"_id": lID, "user_id": uID}
	var animeList AnimeList
	err = collection.FindOne(context.TODO(), filter).Decode(&animeList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("list not found or user is not the creator")
		}
		return nil, fmt.Errorf("could not retrieve the list: %w", err)
	}

	found, err := GetAnimeById(animeId)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$addToSet": bson.M{"anime_list": found.ID},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("could not update the list: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no list found to update")
	}

	fmt.Printf("Successfully updated list with anime ID %v\n", animeId)
	return result, nil

}

func AddCharacterToList(listId string, userId string, characterId string) (interface{}, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	uID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (UserId): %w", err)
	}

	lID, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (ListId): %w", err)
	}

	filter := bson.M{"_id": lID, "user_id": uID}
	var characterList CharacterList
	err = collection.FindOne(context.TODO(), filter).Decode(&characterList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("list not found or user is not the creator")
		}
		return nil, fmt.Errorf("could not retrieve the list: %w", err)
	}

	found, err := GetCharacterById(characterId)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$addToSet": bson.M{"character_list": found.ID},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("could not update the list: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no list found to update")
	}

	fmt.Printf("Successfully updated list with character ID %v\n", characterId)
	return result, nil

}

func GetAllAnimeLists() ([]AnimeList, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var result []AnimeList
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all anime lists")
	return result, nil
}

func GetAllCharacterLists() ([]CharacterList, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var result []CharacterList
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all character lists")
	return result, nil
}

func GetAnimeListById(id string) (*AnimeList, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var result AnimeList
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no anime list found with the given ID")
		}
		return nil, fmt.Errorf("error finding anime list: %w", err)
	}

	// Return the found anime list
	return &result, nil
}

func GetCharacterListById(id string) (*CharacterList, error) {
	client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var result CharacterList
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no character list found with the given ID")
		}
		return nil, fmt.Errorf("error finding character list: %w", err)
	}

	// Return the found character list
	return &result, nil
}
