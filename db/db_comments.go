package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UploadComment(comment Comment, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")
	comment.ID = primitive.NewObjectID()
	comment.Timestamp = time.Now()
	insertResult, err := collection.InsertOne(context.TODO(), comment)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedID)
	return insertResult.InsertedID, nil
}

func GetAllComments(client *mongo.Client) ([]Comment, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var result []Comment
	if err := cursor.All(context.TODO(), &result); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Retrieved all comments")
	return result, nil
}

func GetAllByTypeComments(content_type string, client *mongo.Client) ([]Comment, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

	c_type := CommentType(content_type)

	filter := bson.M{"type": c_type}

	var result []Comment
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no comment found with the given type")
		}
		return nil, fmt.Errorf("error finding comments: %w", err)
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, fmt.Errorf("error decoding comments: %w", err)
	}
	// Return the found comments
	return result, nil
}

func GetCommentById(id string, client *mongo.Client) (*Comment, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}

	var result Comment
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no comment found with the given ID")
		}
		return nil, fmt.Errorf("error finding comment: %w", err)
	}

	// Return the found comment
	return &result, nil
}

func DeleteComment(id string, user_id string, user_role string, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

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

	filter := bson.M{"_id": objID}
	commentData, err := GetCommentById(id, client)
	if err != nil {
		return nil, err
	}

	if commentData.User.UserID != objUserID {
		if user_role != "admin" {
			return nil, fmt.Errorf("only user whom created comment or admins can delete it")
		}
	}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete comment: %w", err)
	}

	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("no comment found with the given ID")
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.DeletedCount)
	return result, nil
}

func UpdateComment(id string, user_id string, text string, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

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
	commentData, err := GetCommentById(id, client)
	if err != nil {
		return nil, err
	}

	if commentData.User.UserID != objUserID {
		return nil, fmt.Errorf("only user whom created comment can edit it")
	}
	update := bson.M{
		"$set": bson.M{
			"text": text,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)

	if err != nil {
		return nil, fmt.Errorf("could not delete comment: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no comment found with the given ID")
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.MatchedCount)
	return result, nil
}

// function to delete comments which content was deleted
func DeleteCommentByContentId(content_id string, content_type string, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(content_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"type": content_type, "content_id": objID}

	result, err := collection.DeleteMany(context.TODO(), filter)

	if err != nil {
		return nil, fmt.Errorf("could not delete comments: %w", err)
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.DeletedCount)
	return result, nil
}

func UpdateCommentRating(id string, value float64, client *mongo.Client) (interface{}, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

	// Convert the string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$inc": bson.M{"rating": value}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, fmt.Errorf("could not update comment rating: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no comment found with the given ID")
	}

	fmt.Printf("Successfully updated %v document(s)\n", result.ModifiedCount)
	return result, nil
}

func GetAllCommentsForContent(content_type string, content_id string, client *mongo.Client) ([]Comment, error) {
	// client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

	c_type := CommentType(content_type)
	c_id, err := primitive.ObjectIDFromHex(content_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"type": c_type, "content_id": c_id}

	var result []Comment
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no comment found with the given type and id")
		}
		return nil, fmt.Errorf("error finding comments: %w", err)
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, fmt.Errorf("error decoding comments: %w", err)
	}
	// Return the found comments
	return result, nil
}

func GetAllCommentsForUser(user_id string, client *mongo.Client) ([]Comment, error) {
	//client := RunMongo()
	collection := client.Database("Anime-Zone").Collection("Comments")

	c_id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter := bson.M{"user.user_id": c_id}

	var result []Comment
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no comment found with the given user")
		}
		return nil, fmt.Errorf("error finding comments: %w", err)
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, fmt.Errorf("error decoding comments: %w", err)
	}
	// Return the found comments
	return result, nil
}
