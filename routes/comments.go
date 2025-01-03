package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostComment(c *gin.Context, client *mongo.Client) {
	var newComment database.Comment
	var newCommentUploader database.CommentUploader

	userIdHex, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	userId, err := primitive.ObjectIDFromHex(userIdHex.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UserId format"})
		return
	}
	username, exists := c.Get("username")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		c.Abort()
		return
	}

	if err := c.BindJSON(&newCommentUploader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime list title"})
		return
	}

	if newCommentUploader.Type == "anime" {
		if !funcs.CheckAnimeExistsById(newCommentUploader.ContentID) {
			// If any anime doesn't exist, return an error message
			c.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime doesn't exist in our db: " + newCommentUploader.ContentID})
			return
		}
	}
	if newCommentUploader.Type == "character" {
		if !funcs.CheckCharactersExistsById(newCommentUploader.ContentID) {
			// If any anime doesn't exist, return an error message
			c.JSON(http.StatusBadRequest, gin.H{"error": "Such Character doesn't exist in our db: " + newCommentUploader.ContentID})
			return
		}
	}
	if newCommentUploader.Type == "anime_list" {
		if !funcs.CheckAnimeListExistsById(newCommentUploader.ContentID) {
			// If any anime doesn't exist, return an error message
			c.JSON(http.StatusBadRequest, gin.H{"error": "Such Anime List doesn't exist in our db: " + newCommentUploader.ContentID})
			return
		}
	}
	if newCommentUploader.Type == "character_list" {
		if !funcs.CheckCharacterListExistsById(newCommentUploader.ContentID) {
			// If any anime doesn't exist, return an error message
			c.JSON(http.StatusBadRequest, gin.H{"error": "Such Character List doesn't exist in our db: " + newCommentUploader.ContentID})
			return
		}
	}
	var commentUser = database.CommentUser{
		UserID:   userId,
		Username: username.(string),
	}
	contentId, err := primitive.ObjectIDFromHex(newCommentUploader.ContentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ContentID format"})
		return
	}
	newComment = database.Comment{
		Type:      newCommentUploader.Type,
		ContentID: contentId,
		User:      commentUser,
		Text:      newCommentUploader.Text,
		Rating:    0,
	}

	insertedID, err := database.UploadComment(newComment, client)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Next Comment Created: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

func GetAllComments(c *gin.Context, client *mongo.Client) {
	comments, err := database.GetAllComments(client)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}
	c.IndentedJSON(http.StatusOK, comments)
}

func GetCommentByType(c *gin.Context, client *mongo.Client) {
	content_type := c.Param("type")

	comment, err := database.GetAllByTypeComments(content_type, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, comment)
}

func GetCommentById(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	comment, err := database.GetCommentById(id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	role, exists := c.Get("role")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Role not found"})
		c.Abort()
		return
	}
	user_id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}

	result, err := database.DeleteComment(id, user_id.(string), role.(string), client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

type UpdateCommentText struct {
	Text string `bson:"text" json:"text"`
}

func UpdateComment(c *gin.Context, client *mongo.Client) {
	var newText UpdateCommentText
	id := c.Param("id")
	user_id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	if err := c.BindJSON(&newText); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text"})
		return
	}
	result, err := database.UpdateComment(id, user_id.(string), newText.Text, client)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}

type UpdateRatingAction struct {
	Action string `json:"action" form:"action"` // "increment" or "decrement"
}

func UpdateCommentRating(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")

	user_id, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	username, exists := c.Get("username")
	if !exists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User Id not found"})
		c.Abort()
		return
	}
	var updateData UpdateRatingAction

	// Bind the action (increment or decrement) from the request
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Determine the increment value based on the action
	var incrementValue int
	if updateData.Action == "increment" {
		incrementValue = 1
	} else if updateData.Action == "decrement" {
		incrementValue = -1
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid action. Use 'increment' or 'decrement'."})
		return
	}

	result, err := database.UpdateCommentRating(id, incrementValue, username.(string), user_id.(string), client)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(result)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Comment rating updated successfully"})
}

func GetCommentForContent(c *gin.Context, client *mongo.Client) {
	content_type := c.Param("type")
	content_id := c.Param("id")

	comment, err := database.GetAllCommentsForContent(content_type, content_id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, comment)
}

func GetCommentForUser(c *gin.Context, client *mongo.Client) {
	user_id := c.Param("id")

	comment, err := database.GetAllCommentsForUser(user_id, client)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, comment)
}
