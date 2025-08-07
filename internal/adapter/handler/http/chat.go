package http

import "main/internal/core/port"

type ChatHandler struct {
	svc port.ChatService
}

func NewChatHandler(svc port.ChatService) *ChatHandler {
	return &ChatHandler{
		svc,
	}
}
