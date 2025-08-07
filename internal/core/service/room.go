package service

import "main/internal/core/port"

type RoomService struct {
	roomRepo port.RoomRepository
}

func NewRoomService(
	roomRepo port.RoomRepository,
) *RoomService {
	return &RoomService{
		roomRepo,
	}
}
