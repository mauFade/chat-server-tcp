package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Message struct {
	ID           bson.ObjectID `bson:"_id"`
	Content      string        `bson:"content"`
	Room         string        `bson:"room"`
	UserNickname string        `bson:"user_nickname"`
	OriginIP     string        `bson:"origin_ip"`
	CreatedAt    time.Time     `bson:"created_at"`
}
