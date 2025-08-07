package http

import (
	"main/internal/core/port"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	svc port.RoomService
}

func NewRoomHandler(svc port.RoomService) *RoomHandler {
	return &RoomHandler{
		svc,
	}
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}

func (h *RoomHandler) DeleteRoomByID(c *gin.Context) {
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}
