package database

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define Type, Status, and ESRB as string types with specific values
type Type string
type Status string
type ESRB string

// Define constants for Type
const (
	Movie Type = "movie"
	TV    Type = "tv"
	OVA   Type = "ova"
	ONA   Type = "ona"
	None  Type = "none"
)

// Define constants for Status
const (
	Finished   Status = "finished"
	Airing     Status = "airing"
	Announced  Status = "announced"
	StatusNone Status = "none"
)

// Define constants for ESRB
const (
	E        ESRB = "E"
	E10Plus  ESRB = "E10+"
	T        ESRB = "T"
	M        ESRB = "M"
	AO       ESRB = "AO"
	RP       ESRB = "RP"
	RP17     ESRB = "RP17+"
	ESRBNone ESRB = "none"
)

type Media struct {
	MediaType string
	MediaLink string
}

// CommentUser struct represents CommentUser information
type RatingUser struct {
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	Username string             `bson:"username" json:"username"`
}

type Rating struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AnimeID   primitive.ObjectID `bson:"anime_id" json:"anime_id"`                 // Reference to the Anime
	User      RatingUser         `bson:"user" json:"user"`                         // ID of the user giving the rating
	Score     float64            `bson:"score" json:"score"`                       // Rating score
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`               // When the rating was given
	Review    string             `bson:"review,omitempty" json:"review,omitempty"` // Optional user review
}

// Define the Anime struct for BSON compatibility and form binding
type Anime struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id" form:"id"` // MongoDB ObjectID
	Title         string             `bson:"title" json:"title" form:"title"`
	ReleaseDate   time.Time          `bson:"release_date" json:"release_date" form:"release_date"` // ISO date string
	AverageRating float64            `bson:"average_rating" json:"average_rating"`                 // Average rating score
	RatingCount   int                `bson:"rating_count" json:"rating_count"`                     // Total number of ratings
	Genre         []string           `bson:"genre" json:"genre" form:"genre"`
	Type          Type               `bson:"type" json:"type" form:"type"`
	Episodes      int                `bson:"episodes" json:"episodes" form:"episodes"`
	Description   string             `bson:"description" json:"description" form:"description"`
	Studio        []string           `bson:"studio" json:"studio" form:"studio"`
	Duration      int                `bson:"duration" json:"duration" form:"duration"` // in minutes
	Status        Status             `bson:"status" json:"status" form:"status"`
	ESRB          ESRB               `bson:"esrb" json:"esrb" form:"esrb"`
	Logo          string             `bson:"logo" json:"logo" form:"logo"`
	Media         []string           `bson:"media" json:"media" form:"media"`
}

type AnimeUploader struct {
	ID            primitive.ObjectID      `bson:"_id,omitempty" json:"id" form:"id"` // MongoDB ObjectID
	Title         string                  `bson:"title" json:"title" form:"title"`
	ReleaseDate   time.Time               `bson:"release_date" json:"release_date" form:"release_date"` // ISO date string
	AverageRating float64                 `bson:"average_rating" json:"average_rating"`                 // Average rating score
	RatingCount   int                     `bson:"rating_count" json:"rating_count"`                     // Total number of ratings
	Genre         []string                `bson:"genre" json:"genre" form:"genre"`
	Type          Type                    `bson:"type" json:"type" form:"type"`
	Episodes      int                     `bson:"episodes" json:"episodes" form:"episodes"`
	Description   string                  `bson:"description" json:"description" form:"description"`
	Studio        []string                `bson:"studio" json:"studio" form:"studio"`
	Duration      int                     `bson:"duration" json:"duration" form:"duration"` // in minutes
	Status        Status                  `bson:"status" json:"status" form:"status"`
	ESRB          ESRB                    `bson:"esrb" json:"esrb" form:"esrb"`
	Link          string                  `bson:"link" json:"link" form:"link"`
	Logo          *multipart.FileHeader   `bson:"logo" json:"logo" form:"logo"`
	Media         []*multipart.FileHeader `bson:"media" json:"media" form:"media"`
}

// Filters struct represents the filters used to search for anime
type Filters struct {
	Genre       []string `bson:"genre" json:"genre"`
	Studio      []string `bson:"studio" json:"studio"`
	ReleaseType Type     `bson:"release_type" json:"release_type"`
	Status      Status   `bson:"status" json:"status"`
	ESRB        ESRB     `bson:"esrb" json:"esrb"`
}

type FromAnime struct {
	ID    primitive.ObjectID `bson:"id" json:"id"`
	Title string             `bson:"title" json:"title"`
}

// Character struct represents character information
type Character struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Age       int                `bson:"age" json:"age"`
	FromAnime []FromAnime        `bson:"from_anime" json:"from_anime"`
	Gender    string             `bson:"gender" json:"gender"`
	Bio       string             `bson:"bio" json:"bio"`
	Status    string             `bson:"status" json:"status"` // e.g., "alive", "dead", "unknown"
	Logo      string             `bson:"logo" json:"logo"`
	Media     []string           `bson:"media" json:"media"`
}

type CharacterUploader struct {
	ID        primitive.ObjectID      `bson:"_id" json:"id" form:"id"`
	FirstName string                  `bson:"first_name" json:"first_name" form:"first_name"`
	LastName  string                  `bson:"last_name" json:"last_name" form:"last_name"`
	Age       int                     `bson:"age" json:"age" form:"age"`
	FromAnime []string                `bson:"from_anime" json:"from_anime" form:"from_anime"`
	Gender    string                  `bson:"gender" json:"gender" form:"gender"`
	Bio       string                  `bson:"bio" json:"bio" form:"bio"`
	Status    string                  `bson:"status" json:"status" form:"status"` // e.g., "alive", "dead", "unknown"
	Logo      *multipart.FileHeader   `bson:"logo" json:"logo" form:"logo"`
	Media     []*multipart.FileHeader `bson:"media" json:"media" form:"media"`
}

// User struct represents User information
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Role     string             `bson:"role" json:"role"`
	Bio      string             `bson:"bio" json:"bio"`
	Logo     string             `bson:"logo" json:"logo"`
}

// User struct represents User information
type UserUploader struct {
	ID       primitive.ObjectID    `bson:"_id" json:"id" form:"id"`
	Email    string                `bson:"email" json:"email" form:"email"`
	Username string                `bson:"username" json:"username" form:"username"`
	Password string                `bson:"password" json:"password" form:"password"`
	Role     string                `bson:"role" json:"role" form:"role"`
	Bio      string                `bson:"bio" json:"bio" form:"bio"`
	Logo     *multipart.FileHeader `bson:"logo" json:"logo" form:"logo"`
}

// AnimeList struct represents AnimeList information
type AnimeList struct {
	ID        primitive.ObjectID   `bson:"_id" json:"id"`
	Name      string               `bson:"name" json:"name"`
	UserID    primitive.ObjectID   `bson:"user_id" json:"user_id"`
	AnimeList []primitive.ObjectID `bson:"anime_list" json:"anime_list"`
	Rating    float64              `bson:"rating" json:"rating"`
}

// CharacterList struct represents CharacterList information
type CharacterList struct {
	ID            primitive.ObjectID   `bson:"_id" json:"id"`
	Name          string               `bson:"name" json:"name"`
	UserID        primitive.ObjectID   `bson:"user_id" json:"user_id"`
	CharacterList []primitive.ObjectID `bson:"character_list" json:"character_list"`
	Rating        float64              `bson:"rating" json:"rating"`
}

// CommentUser struct represents CommentUser information
type CommentUser struct {
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	Username string             `bson:"username" json:"username"`
}
type CommentType string

// Define constants for Status
const (
	TypeAnime         CommentType = "anime"
	TypeCharacter     CommentType = "character"
	TypeAnimeList     CommentType = "anime_list"
	TypeCharacterList CommentType = "character_list"
)

// Comment struct represents Comment information
type Comment struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Type      CommentType        `bson:"type" json:"type"`
	ContentID primitive.ObjectID `bson:"content_id" json:"content_id"`
	User      CommentUser        `bson:"user" json:"user"`
	Text      string             `bson:"text" json:"text"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp" form:"timestamp"`
	Rating    float64            `bson:"rating" json:"rating"`
}
type CommentUploader struct {
	Type      CommentType `bson:"type" json:"type"`
	ContentID string      `bson:"content_id" json:"content_id"`
	User      CommentUser `bson:"user" json:"user"`
	Text      string      `bson:"text" json:"text"`
	Rating    float64     `bson:"rating" json:"rating"`
}
