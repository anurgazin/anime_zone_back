package database

import "time"

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

// Anime struct represents the anime information
type Anime struct {
	ID          string    `json:"id"` // UUID format
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"` // ISO date string
	Rating      float64   `json:"rating"`
	Genre       []string  `json:"genre"`
	Type        Type      `json:"type"`
	Episodes    int       `json:"episodes"`
	Description string    `json:"description"`
	Studio      []string  `json:"studio"`
	Duration    int       `json:"duration"` // in minutes
	Status      Status    `json:"status"`
	ESRB        ESRB      `json:"esrb"`
}

// Filters struct represents the filters used to search for anime
type Filters struct {
	Genre       []string `json:"genre"`
	Studio      []string `json:"studio"`
	ReleaseType Type     `json:"release_type"`
	Status      Status   `json:"status"`
	ESRB        ESRB     `json:"esrb"`
}

// Character struct represents character information
type Character struct {
	ID        string   `json:"id"` // UUID format
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	FromAnime []string `json:"from_anime"`
	Gender    string   `json:"gender"`
	Bio       string   `json:"bio"`
	Status    string   `json:"status"` // e.g., "alive", "dead", "unknown"
}
