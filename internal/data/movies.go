package data

import (
	"time"
)

// !!! All the fields in our Movie struct are exported (i.e. start with a capital letter),
// !!! which is necessary for them to be visible to Go’s encoding/json package.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Hyphen directive hides it
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty,string"` //force to be string in json response
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
