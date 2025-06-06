package repository

import (
	"context"
	"time"

	"github.com/mauFade/chat-server-tcp/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (r *MessageRepository) FindByRoom(room string) ([]*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var messages []*models.Message

	opts := options.Find().SetLimit(20)
	filter := bson.D{
		{Key: "room", Value: room},
	}

	data, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = data.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
