package routes

import (
	database "anime_zone/back_end/db"
	"anime_zone/back_end/funcs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostComment(c *gin.Context) {
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
		Rating:    newCommentUploader.Rating,
	}

	insertedID, err := database.UploadComment(newComment)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Next Comment Created: %v", insertedID)
	c.IndentedJSON(http.StatusCreated, result)
}

func GetAllComments(c *gin.Context) {
	comments, err := database.GetAllComments()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}
	c.IndentedJSON(http.StatusOK, comments)
}

func GetCommentByType(c *gin.Context) {
	content_type := c.Param("type")

	comment, err := database.GetAllByTypeComments(content_type)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, comment)
}