package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	RoomID   primitive.ObjectID `json:"room_id" bson:"room_id"`
	Message  string             `json:"message" bson:"message"`
	Type     string             `json:"type" bson:"type"` // chat, reply, image
	UpdateAt time.Time          `json:"update_at" bson:"update_at"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
}
