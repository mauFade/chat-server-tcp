package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id"`
	Nickname  string        `bson:"nickname"`
	Room      string        `bson:"current_room"`
	LastIP    string        `bson:"last_ip"`
	CreatedAt time.Time     `bson:"created_at"`
}
