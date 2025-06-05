package models

import "time"

type Room struct {
	ID          string    `bson:"id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
}
