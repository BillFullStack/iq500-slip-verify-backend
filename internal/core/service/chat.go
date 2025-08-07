package service

import (
	"main/internal/core/port"
)

type ChatService struct {
	chatRepo port.ChatRepository
}

func NewChatService(
	chatRepo port.ChatRepository,
) *ChatService {
	return &ChatService{
		chatRepo,
	}
}
