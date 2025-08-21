package repository

import (
	"context"
	"fmt"
	"main/internal/core/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ROOM_COLLECTION = "room"
)

type RoomRepository struct {
	DB *mongo.Database
}

func NewRoomRepository(DB *mongo.Database) *RoomRepository {
	return &RoomRepository{
		DB,
	}
}

func (r *RoomRepository) GetRoomByUserId(id primitive.ObjectID) ([]domain.Room, error) {
	c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	option := options.Find()
	option.SetSort(bson.M{
		"update_at": -1,
	})
	defer cancel()

	var room []domain.Room

	filter := bson.M{
		"user_id": id,
	}

	obj, err := r.DB.Collection(ROOM_COLLECTION).Find(c, filter, option)
	if err != nil {
		return nil, err
	}

	err = obj.All(c, &room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *RoomRepository) CreateRoom(payload domain.Room) error {
	c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.DB.Collection(ROOM_COLLECTION).InsertOne(c, payload)

	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) DeleteRoomByID(id primitive.ObjectID) error {
	c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	_, err := r.DB.Collection(ROOM_COLLECTION).DeleteOne(c, filter)
	if err != nil {
		fmt.Println("delete category platform by business code err: ", err)
		return err
	}
	return nil
}
