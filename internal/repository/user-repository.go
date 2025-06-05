package repository

import (
	"context"
	"time"

	"github.com/mauFade/chat-server-tcp/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func (r *UserRepository) FindByNickname(nick string) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "nickname", Value: nick},
	}

	var u *models.User

	_ = r.collection.FindOne(ctx, filter).Decode(&u)

	return u
}
