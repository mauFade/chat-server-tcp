package repository

import (
	"context"
	"time"

	"github.com/mauFade/chat-server-tcp/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RoomRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRoomRepository(client *mongo.Client) *RoomRepository {
	return &RoomRepository{
		client:     client,
		collection: client.Database("chat-server-tcp-db").Collection("rooms"),
	}
}

func (r *RoomRepository) CreateRoom(room models.Room) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, room)
	return err
}

func (r *RoomRepository) FindManyRooms() []*models.Room {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
	}
	// Unpacks the cursor into a slice
	var results []*models.Room
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return results
}
