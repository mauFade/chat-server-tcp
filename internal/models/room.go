package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Room struct {
	ID          bson.ObjectID `bson:"_id"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	CreatedAt   time.Time     `bson:"created_at"`
}
