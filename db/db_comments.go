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
