package repository

import "go.mongodb.org/mongo-driver/mongo"

const (
	CHAT_COLLECTION = "chat"
)

type ChatRepository struct {
	DB *mongo.Database
}

func NewChatRepository(DB *mongo.Database) *ChatRepository {
	return &ChatRepository{
		DB,
	}
}
