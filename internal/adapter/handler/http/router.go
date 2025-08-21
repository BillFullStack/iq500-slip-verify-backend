package http

import (
	"main/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	resource *mongo.Database,
	chatHandler *ChatHandler,
	roomHandler *RoomHandler,
	authHandler *AuthenticationHandler,
) (*Router, error) {
	router := gin.Default()
	router.Use(CORS)

	// ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"version": "1.0.1",
		})
	})

	api := router.Group("/api")
	auth := router.Group("/api")
	auth.Use(utils.Authentication(resource))

	// authentication
	api.POST("/login", authHandler.Login)
	api.POST("/register", authHandler.Register)
	api.POST("/reset-password", authHandler.ResetPassword)

	// chat
	chat := auth.Group("/chat")
	{
		chat.GET("/:id", chatHandler.GetChatByRoomID)
		chat.POST("", chatHandler.Chat)
	}

	// room
	room := auth.Group("/room")
	{
		room.GET("", roomHandler.GetRoom)
		room.POST("", roomHandler.CreateRoom)
		room.DELETE("/:id", roomHandler.DeleteRoomByID)
	}

	return &Router{
		router,
	}, nil
}

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
		return
	} else {
		c.AbortWithStatus(http.StatusOK)
		return
	}
}

func (r *Router) Serve(listenAddr string) error {
	srv := &http.Server{
		Addr:    listenAddr,
		Handler: r,
	}
	return srv.ListenAndServe()
}
