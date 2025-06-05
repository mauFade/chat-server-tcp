package repository

import (
	"context"
	"time"

	"github.com/mauFade/chat-server-tcp/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client:     client,
		collection: client.Database("chat-server-tcp-db").Collection("users"),
	}
}

func (r *UserRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}
