package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  primitive.ObjectID `json:"username" bson:"username"`
	Password  primitive.ObjectID `json:"password" bson:"password"`
	UpdateAt  time.Time          `json:"update_at" bson:"update_at"`
	CreateAt  time.Time          `json:"create_at" bson:"create_at"`
	LastLogin time.Time          `json:"last_login" bson:"last_login"`
}
