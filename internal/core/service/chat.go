package service

import (
	"errors"
	"fmt"
	"main/internal/core/domain"
	"main/internal/core/port"
	"main/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *ChatService) GetChatByRoomID(c *gin.Context, id primitive.ObjectID) ([]domain.Chat, error) {
	chat, err := s.chatRepo.GetChatByRoomID(id)
	if err != nil {
		fmt.Println("error get chat", err)
		utils.Response(c, http.StatusBadRequest, 1, "เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง", err.Error(), nil)
		return nil, err
	}

	return chat, nil
}

func (s *ChatService) Chat(c *gin.Context, payload domain.PayloadChat) error {
	if payload.Message == "" && payload.Img == "" {
		fmt.Println("err message and img is empty")
		utils.Response(c, http.StatusBadRequest, 1, "เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง", "err message and img is empty", nil)
		return errors.New("err message and img is empty")
	}

	messageType := ""
	if payload.Message != "" {
		messageType = "text"
	} else {
		messageType = "image"
	}

	roomID, _ := primitive.ObjectIDFromHex(payload.RoomID)

	chat := domain.Chat{
		ID:          primitive.NewObjectID(),
		RoomID:      roomID,
		Type:        "reply",
		MessageType: messageType,
		Message:     payload.Message,
		UpdateAt:    time.Now(),
		CreateAt:    time.Now(),
	}

	err := s.chatRepo.CreateChat(chat)
	if err != nil {
		fmt.Println("error create chat", err)
		utils.Response(c, http.StatusBadRequest, 1, "เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง", err.Error(), nil)
		return err
	}

	return nil
}
