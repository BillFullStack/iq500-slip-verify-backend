package repository

import "go.mongodb.org/mongo-driver/mongo"

const (
	USER_COLLECTION = "user"
)

type UserRepository struct {
	DB *mongo.Database
}

func NewUserRepository(DB *mongo.Database) *UserRepository {
	return &UserRepository{
		DB,
	}
}
