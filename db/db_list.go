package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAnimeList(list PostListRequest, client *mongo.Client) (interface{}, error) {
	var animeList AnimeList
	var listUser ListUser
	collection := client.Database("Anime-Zone").Collection("AnimeList")
	animeList.ID = primitive.NewObjectID()
	userId, err := primitive.ObjectIDFromHex(list.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UserId format: %w", err)
	}
	listUser.UserID = userId
	listUser.Username = list.Username
	animeList.User = listUser
	animeList.Name = list.ListTitle
	animeList.Public = list.Public
	animeList.AnimeList = []primitive.ObjectID{}
	for _, a := range list.ContentList {
		animeId, err := primitive.ObjectIDFromHex(a)
		if err != nil {
			return nil, fmt.Errorf("invalid AnimeId format: %w", err)
		}
		animeList.AnimeList = append(animeList.AnimeList, animeId)
	}
	insertResult, err := collection.InsertOne(context.TODO(), animeList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func CreateCharacterList(list PostListRequest, client *mongo.Client) (interface{}, error) {
	var characterList CharacterList
	var listUser ListUser
	collection := client.Database("Anime-Zone").Collection("CharacterList")
	characterList.ID = primitive.NewObjectID()
	userId, err := primitive.ObjectIDFromHex(list.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UserId format: %w", err)
	}
	listUser.UserID = userId
	listUser.Username = list.Username
	characterList.User = listUser
	characterList.Name = list.ListTitle
	characterList.Public = list.Public
	characterList.CharacterList = []primitive.ObjectID{}
	for _, c := range list.ContentList {
		characterId, err := primitive.ObjectIDFromHex(c)
		if err != nil {
			return nil, fmt.Errorf("invalid AnimeId format: %w", err)
		}
		characterList.CharacterList = append(characterList.CharacterList, characterId)
	}
	insertResult, err := collection.InsertOne(context.TODO(), characterList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func AddAnimeToList(listId string, userId string, animeId string, client *mongo.Client) (interface{}, error) {
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	uID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (UserId): %w", err)
	}

	lID, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (ListId): %w", err)
	}

	filter := bson.M{"_id": lID, "user.user_id": uID}
	var animeList AnimeList
	err = collection.FindOne(context.TODO(), filter).Decode(&animeList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("list not found or user is not the creator")
		}
		return nil, fmt.Errorf("could not retrieve the list: %w", err)
	}

	found, err := GetAnimeById(animeId, client)
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

func AddCharacterToList(listId string, userId string, characterId string, client *mongo.Client) (interface{}, error) {
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	uID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (UserId): %w", err)
	}

	lID, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format (ListId): %w", err)
	}

	filter := bson.M{"_id": lID, "user.user_id": uID}
	var characterList CharacterList
	err = collection.FindOne(context.TODO(), filter).Decode(&characterList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("list not found or user is not the creator")
		}
		return nil, fmt.Errorf("could not retrieve the list: %w", err)
	}

	found, err := GetCharacterById(characterId, client)
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

func GetAllAnimeLists(client *mongo.Client) ([]AnimeList, error) {
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

func GetAllAnimeListsToDisplay(client *mongo.Client) ([]AnimeList, error) {
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	cursor, err := collection.Find(context.TODO(), bson.M{"public": true})
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

func GetAllCharacterLists(client *mongo.Client) ([]CharacterList, error) {
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
func GetAllCharacterListsToDisplay(client *mongo.Client) ([]CharacterList, error) {
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	cursor, err := collection.Find(context.TODO(), bson.M{"public": true})
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

func GetAnimeListById(id string, client *mongo.Client) (*AnimeList, error) {
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

func GetCharacterListById(id string, client *mongo.Client) (*CharacterList, error) {
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

func UpdateAnimeList(id string, user_id string, text string, content []string, public bool, client *mongo.Client) (interface{}, error) {
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	// Convert the string ID to ObjectID
	objUserID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}
	listData, err := GetAnimeListById(id, client)
	if err != nil {
		return nil, err
	}

	if listData.User.UserID != objUserID {
		return nil, fmt.Errorf("only user whom created list can edit it")
	}

	animeList := []primitive.ObjectID{}
	for _, a := range content {
		animeId, err := primitive.ObjectIDFromHex(a)
		if err != nil {
			return nil, fmt.Errorf("invalid AnimeId format: %w", err)
		}
		animeList = append(animeList, animeId)
	}

	update := bson.M{
		"$set": bson.M{
			"name":       text,
			"anime_list": animeList,
			"public":     public,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)

	if err != nil {
		return nil, fmt.Errorf("could not update list: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no list found with the given ID")
	}

	fmt.Printf("Successfully updated %v list(s)\n", result.MatchedCount)
	return result, nil
}

func UpdateCharacterList(id string, user_id string, text string, content []string, public bool, client *mongo.Client) (interface{}, error) {
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	// Convert the string ID to ObjectID
	objUserID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}
	listData, err := GetCharacterListById(id, client)
	if err != nil {
		return nil, err
	}

	if listData.User.UserID != objUserID {
		return nil, fmt.Errorf("only user whom created list can edit it")
	}

	characterList := []primitive.ObjectID{}
	for _, c := range content {
		characterId, err := primitive.ObjectIDFromHex(c)
		if err != nil {
			return nil, fmt.Errorf("invalid AnimeId format: %w", err)
		}
		characterList = append(characterList, characterId)
	}

	update := bson.M{
		"$set": bson.M{
			"name":           text,
			"character_list": characterList,
			"public":         public,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)

	if err != nil {
		return nil, fmt.Errorf("could not update list: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no list found with the given ID")
	}

	fmt.Printf("Successfully updated %v list(s)\n", result.MatchedCount)
	return result, nil
}

func UpdateListRating(id, listType, userID, username string, value int, client *mongo.Client) (interface{}, error) {
	scoreCollection := client.Database("Anime-Zone").Collection("Score")

	// Convert IDs
	listID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid list ID: %w", err)
	}
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// Determine collection and filter based on listType
	var (
		listCollection *mongo.Collection
		scoreType      ScoreType
	)
	switch listType {
	case "anime_list":
		listCollection = client.Database("Anime-Zone").Collection("AnimeList")
		scoreType = ScoreTypeAnimeList
	case "character_list":
		listCollection = client.Database("Anime-Zone").Collection("CharacterList")
		scoreType = ScoreTypeCharacterList
	default:
		return nil, fmt.Errorf("invalid list type")
	}

	filter := bson.M{
		"content_type": scoreType,
		"content_id":   listID,
		"user.user_id": userObjectID,
	}

	// Check for existing score
	var existingScore Score
	err = scoreCollection.FindOne(context.TODO(), filter).Decode(&existingScore)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("error finding existing score: %w", err)
	}

	if err == nil {
		// Update existing score
		update := bson.M{
			"$set": bson.M{"score": value, "timestamp": time.Now()},
		}
		if _, err := scoreCollection.UpdateOne(context.TODO(), filter, update); err != nil {
			return nil, fmt.Errorf("error updating score: %w", err)
		}
	} else {
		// Insert new score
		newScore := Score{
			ID:          primitive.NewObjectID(),
			ContentID:   listID,
			User:        ScoreUser{UserID: userObjectID, Username: username},
			Score:       value,
			Timestamp:   time.Now(),
			ContentType: scoreType,
		}
		if _, err := scoreCollection.InsertOne(context.TODO(), newScore); err != nil {
			return nil, fmt.Errorf("error inserting score: %w", err)
		}
	}

	// Update the list rating
	if err := updateListRating(listID, scoreType, scoreCollection, listCollection); err != nil {
		return nil, fmt.Errorf("error updating list rating: %w", err)
	}

	return "Score updated successfully", nil
}

func updateListRating(listID primitive.ObjectID, scoreType ScoreType, scoreCollection, listCollection *mongo.Collection) error {
	cursor, err := scoreCollection.Find(context.TODO(), bson.M{"content_type": scoreType, "content_id": listID})
	if err != nil {
		return fmt.Errorf("error finding scores: %w", err)
	}
	defer cursor.Close(context.TODO())

	// Calculate total score
	totalScore := 0
	for cursor.Next(context.TODO()) {
		var score Score
		if err := cursor.Decode(&score); err != nil {
			return fmt.Errorf("error decoding score: %w", err)
		}
		totalScore += score.Score
	}

	// Update list rating
	_, err = listCollection.UpdateOne(context.TODO(), bson.M{"_id": listID}, bson.M{"$set": bson.M{"rating": totalScore}})
	if err != nil {
		return fmt.Errorf("error updating list rating: %w", err)
	}
	return nil
}

func GetAllAnimeListsByAnimeId(anime_id string, client *mongo.Client) ([]AnimeList, error) {
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(anime_id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"anime_list": objID, "public": true}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var result []AnimeList
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all anime lists with given anime")
	return result, nil
}

func GetAllCharacterListsByCharacterId(character_id string, client *mongo.Client) ([]CharacterList, error) {
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(character_id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"character_list": objID, "public": true}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var result []CharacterList
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all characters lists with given character")
	return result, nil
}

func GetAllAnimeListsByUserId(user_id string, client *mongo.Client) ([]AnimeList, error) {
	collection := client.Database("Anime-Zone").Collection("AnimeList")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"user.user_id": objID}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var result []AnimeList
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all anime lists with given user id")
	return result, nil
}

func GetAllCharacterListsByUserId(user_id string, client *mongo.Client) ([]CharacterList, error) {
	collection := client.Database("Anime-Zone").Collection("CharacterList")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"user.user_id": objID}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var result []CharacterList
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all characters lists with given user")
	return result, nil
}

func DeleteAnimeList(id string, user_id string, user_role string, client *mongo.Client) (interface{}, error) {
	anime_list_collection := client.Database("Anime-Zone").Collection("AnimeList")
	//character_collection := client.Database("Anime-Zone").Collection("CharacterList")
	// comment_collection := client.Database("Anime-Zone").Collection("Comment")
	score_collection := client.Database("Anime-Zone").Collection("Score")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	// Convert the string ID to ObjectID
	usrID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var animeList AnimeList

	err = anime_list_collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&animeList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no anime list found with the given ID")
		}
		return nil, fmt.Errorf("could not fetch anime list: %w", err)
	}

	// Check if the user is the author or an admin
	if animeList.User.UserID != usrID && user_role != "admin" {
		return nil, fmt.Errorf("unauthorized: only the author or an admin can delete this anime list")
	}

	anime_result, err := anime_list_collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete anime: %w", err)
	}

	if anime_result.DeletedCount == 0 {
		return nil, fmt.Errorf("no anime found with the given ID")
	}

	score_filter := bson.M{"content_type": "anime_list", "content_id": objID}
	score_result, err := score_collection.DeleteMany(context.TODO(), score_filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete anime list: %w", err)
	}

	comment_result, err := DeleteCommentByContentId(id, "anime_list", client)
	if err != nil {
		return nil, fmt.Errorf("error during deleting comments")
	}

	return map[string]interface{}{
		"deleted_anime_list_count": anime_result.DeletedCount,
		"deleted_scores_count":     score_result.DeletedCount,
		"deleted_comment_count":    comment_result,
	}, nil
}

func DeleteCharactersList(id string, user_id string, user_role string, client *mongo.Client) (interface{}, error) {
	character_list_collection := client.Database("Anime-Zone").Collection("CharacterList")
	//character_collection := client.Database("Anime-Zone").Collection("CharacterList")
	// comment_collection := client.Database("Anime-Zone").Collection("Comment")
	score_collection := client.Database("Anime-Zone").Collection("Score")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	// Convert the string ID to ObjectID
	usrID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var characterList CharacterList

	err = character_list_collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&characterList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no character list found with the given ID")
		}
		return nil, fmt.Errorf("could not fetch character list: %w", err)
	}

	// Check if the user is the author or an admin
	if characterList.User.UserID != usrID && user_role != "admin" {
		return nil, fmt.Errorf("unauthorized: only the author or an admin can delete this character list")
	}

	character_list_result, err := character_list_collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete character list: %w", err)
	}

	if character_list_result.DeletedCount == 0 {
		return nil, fmt.Errorf("no character list found with the given ID")
	}

	score_filter := bson.M{"content_type": "character_list", "content_id": objID}
	score_result, err := score_collection.DeleteMany(context.TODO(), score_filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete character list: %w", err)
	}

	comment_result, err := DeleteCommentByContentId(id, "character_list", client)
	if err != nil {
		return nil, fmt.Errorf("error during deleting comments")
	}

	return map[string]interface{}{
		"deleted_character_list_count": character_list_result.DeletedCount,
		"deleted_scores_count":         score_result.DeletedCount,
		"deleted_comment_count":        comment_result,
	}, nil
}
