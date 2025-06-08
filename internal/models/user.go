package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserStatus string

const (
	StatusOnline  UserStatus = "online"
	StatusAway    UserStatus = "away"
	StatusBusy    UserStatus = "busy"
	StatusOffline UserStatus = "offline"
)

type User struct {
	ID           bson.ObjectID   `bson:"_id,omitempty"`
	Nickname     string          `bson:"nickname"`
	Password     string          `bson:"password,omitempty"`
	Email        string          `bson:"email,omitempty"`
	Room         string          `bson:"room"`
	LastIP       string          `bson:"last_ip"`
	Status       UserStatus      `bson:"status"`
	LastSeen     time.Time       `bson:"last_seen"`
	CreatedAt    time.Time       `bson:"created_at"`
	UpdatedAt    time.Time       `bson:"updated_at"`
	IsAdmin      bool            `bson:"is_admin"`
	Preferences  UserPreferences `bson:"preferences"`
	BlockedUsers []string        `bson:"blocked_users"`
}

type UserPreferences struct {
	Theme            string `bson:"theme"`
	Notifications    bool   `bson:"notifications"`
	SoundEnabled     bool   `bson:"sound_enabled"`
	MessageHistory   int    `bson:"message_history"`
	AutoJoinRoom     string `bson:"auto_join_room"`
	ShowOnlineStatus bool   `bson:"show_online_status"`
}

func NewUser(nickname string, ip string) *User {
	now := time.Now()
	return &User{
		ID:        bson.NewObjectID(),
		Nickname:  nickname,
		LastIP:    ip,
		Status:    StatusOnline,
		LastSeen:  now,
		CreatedAt: now,
		UpdatedAt: now,
		Preferences: UserPreferences{
			Theme:            "default",
			Notifications:    true,
			SoundEnabled:     true,
			MessageHistory:   100,
			ShowOnlineStatus: true,
		},
		BlockedUsers: make([]string, 0),
	}
}
