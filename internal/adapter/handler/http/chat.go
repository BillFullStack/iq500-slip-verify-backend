package http

import (
	"main/internal/core/port"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}

func (h *ChatHandler) Chat(c *gin.Context) {
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}
