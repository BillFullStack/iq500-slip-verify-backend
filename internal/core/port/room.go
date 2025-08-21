package port

import (
	"main/internal/core/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomRepository interface {
	GetRoomByUserId(id primitive.ObjectID) ([]domain.Room, error)
	CreateRoom(payload domain.Room) error
	DeleteRoomByID(id primitive.ObjectID) error
}

type RoomService interface {
	GetRoomByUserId(c *gin.Context, id primitive.ObjectID) ([]domain.Room, error)
	CreateRoom(c *gin.Context, payload domain.PayloadRoom) error
	DeleteRoomByID(c *gin.Context, id primitive.ObjectID) error
}
