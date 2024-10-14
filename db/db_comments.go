package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UploadComment(comment Comment) (interface{}, error) {
	client := RunMongo()
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

func GetAllComments() ([]Comment, error) {
	client := RunMongo()
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

func GetAllByTypeComments(content_type string) ([]Comment, error) {
	client := RunMongo()
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

func GetCommentById(id string) (*Comment, error) {
	client := RunMongo()
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

func DeleteComment(id string, user_id string, user_role string) (interface{}, error) {
	client := RunMongo()
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
	commentData, err := GetCommentById(id)
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
