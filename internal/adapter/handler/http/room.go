package http

import "main/internal/core/port"

type RoomHandler struct {
	svc port.RoomService
}

func NewRoomHandler(svc port.RoomService) *RoomHandler {
	return &RoomHandler{
		svc,
	}
}
