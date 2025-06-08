package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Message struct {
	ID               bson.ObjectID   `bson:"_id,omitempty"`
	Content          string          `bson:"content"`
	FormattedContent string          `bson:"formatted_content"`
	Room             string          `bson:"room"`
	UserNickname     string          `bson:"user_nickname"`
	OriginIP         string          `bson:"origin_ip"`
	CreatedAt        time.Time       `bson:"created_at"`
	Type             MessageType     `bson:"type"`
	Metadata         MessageMetadata `bson:"metadata,omitempty"`
}

type MessageType string

const (
	MessageTypeText   MessageType = "text"
	MessageTypeEmoji  MessageType = "emoji"
	MessageTypeFile   MessageType = "file"
	MessageTypeSystem MessageType = "system"
)

type MessageMetadata struct {
	FileURL  string        `bson:"file_url,omitempty"`
	FileType string        `bson:"file_type,omitempty"`
	FileSize int64         `bson:"file_size,omitempty"`
	IsEdited bool          `bson:"is_edited"`
	EditedAt time.Time     `bson:"edited_at,omitempty"`
	ReplyTo  bson.ObjectID `bson:"reply_to,omitempty"`
}
