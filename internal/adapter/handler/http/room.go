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

type RoomHandler struct {
	svc port.RoomService
}

func NewRoomHandler(svc port.RoomService) *RoomHandler {
	return &RoomHandler{
		svc,
	}
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	fmt.Println("GetRoom")

	id, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))

	room, err := h.svc.GetRoomByUserId(c, id)
	if err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 200, "success", "ok", room)
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	fmt.Println("CreateRoom")

	var payload domain.PayloadRoom
	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println("error bind", err)
		utils.Response(c, http.StatusBadRequest, 1, "ข้อมูลไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง", err.Error(), nil)
		return
	}

	if err := h.svc.CreateRoom(c, payload); err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}

func (h *RoomHandler) DeleteRoomByID(c *gin.Context) {
	fmt.Println("DeleteRoomByID")

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	if err := h.svc.DeleteRoomByID(c, id); err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}
