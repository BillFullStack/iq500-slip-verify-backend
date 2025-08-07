package http

import (
	"main/internal/adapter/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	chatHandler ChatHandler,
	roomHandler RoomHandler,
	userHandler UserHandler,
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
