package database

import (
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

// Define the Anime struct for BSON compatibility
type Anime struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB ObjectID
	Title       string             `bson:"title" json:"title"`
	ReleaseDate time.Time          `bson:"release_date" json:"release_date"` // ISO date string
	Rating      float64            `bson:"rating" json:"rating"`
	Genre       []string           `bson:"genre" json:"genre"`
	Type        Type               `bson:"type" json:"type"`
	Episodes    int                `bson:"episodes" json:"episodes"`
	Description string             `bson:"description" json:"description"`
	Studio      []string           `bson:"studio" json:"studio"`
	Duration    int                `bson:"duration" json:"duration"` // in minutes
	Status      Status             `bson:"status" json:"status"`
	ESRB        ESRB               `bson:"esrb" json:"esrb"`
	Logo        []string           `bson:"logo" json:"logo"`
	Media       []string           `bson:"media" json:"media"`
}

// Filters struct represents the filters used to search for anime
type Filters struct {
	Genre       []string `bson:"genre" json:"genre"`
	Studio      []string `bson:"studio" json:"studio"`
	ReleaseType Type     `bson:"release_type" json:"release_type"`
	Status      Status   `bson:"status" json:"status"`
	ESRB        ESRB     `bson:"esrb" json:"esrb"`
}

// Character struct represents character information
type Character struct {
	ID        primitive.ObjectID   `bson:"_id" json:"id"`
	FirstName string               `bson:"first_name" json:"first_name"`
	LastName  string               `bson:"last_name" json:"last_name"`
	Age       int                  `bson:"age" json:"age"`
	FromAnime []primitive.ObjectID `bson:"from_anime" json:"from_anime"`
	Gender    string               `bson:"gender" json:"gender"`
	Bio       string               `bson:"bio" json:"bio"`
	Status    string               `bson:"status" json:"status"` // e.g., "alive", "dead", "unknown"
	Logo      []string             `bson:"logo" json:"logo"`
	Media     []string             `bson:"media" json:"media"`
}

// Character struct represents character information
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Role     string             `bson:"role" json:"role"`
	Bio      string             `bson:"bio" json:"bio"`
	Logo     []string           `bson:"logo" json:"logo"`
}

type AnimeList struct {
	ID        primitive.ObjectID   `bson:"_id" json:"id"`
	Name      string               `bson:"name" json:"name"`
	UserID    primitive.ObjectID   `bson:"user_id" json:"user_id"`
	AnimeList []primitive.ObjectID `bson:"anime_list" json:"anime_list"`
	Rating    float64              `bson:"rating" json:"rating"`
}
type CharacterList struct {
	ID            primitive.ObjectID   `bson:"_id" json:"id"`
	Name          string               `bson:"name" json:"name"`
	UserID        primitive.ObjectID   `bson:"user_id" json:"user_id"`
	CharacterList []primitive.ObjectID `bson:"anime_list" json:"anime_list"`
	Rating        float64              `bson:"rating" json:"rating"`
}
