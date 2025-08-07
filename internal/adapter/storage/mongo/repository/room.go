package repository

import "go.mongodb.org/mongo-driver/mongo"

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
