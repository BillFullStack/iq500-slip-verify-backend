package http

import (
	"fmt"
	"main/internal/core/domain"
	"main/internal/core/port"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHandler struct {
	svc port.ChatService
}

func NewChatHandler(svc port.ChatService) *ChatHandler {
	return &ChatHandler{
		svc,
	}
}

func (h *ChatHandler) GetChatByRoomID(c *gin.Context) {
	fmt.Println("GetChatByRoomID")

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	chat, err := h.svc.GetChatByRoomID(c, id)
	if err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 0, "success", "ok", chat)
}

func (h *ChatHandler) Chat(c *gin.Context) {
	fmt.Println("Chat")

	var payload domain.PayloadChat
	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println("error bind", err)
		utils.Response(c, http.StatusBadRequest, 1, "ข้อมูลไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง", err.Error(), nil)
		return
	}

	if err := h.svc.Chat(c, payload); err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 0, "success", "ok", nil)
}
