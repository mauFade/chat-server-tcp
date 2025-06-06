package repository

import (
	"context"
	"time"

	"github.com/mauFade/chat-server-tcp/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MessageRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMessageRepository(client *mongo.Client) *MessageRepository {
	return &MessageRepository{
		client:     client,
		collection: client.Database("chat-server-tcp-db").Collection("messages"),
	}
}

func (r *MessageRepository) CreateMessage(m models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, m)
	return err
}
