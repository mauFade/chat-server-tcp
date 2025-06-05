package models

import "time"

type User struct {
	ID        string    `bson:"id"`
	Nickname  string    `bson:"nickname"`
	Room      string    `bson:"current_room"`
	LastIP    string    `bson:"last_ip"`
	CreatedAt time.Time `bson:"created_at"`
}
