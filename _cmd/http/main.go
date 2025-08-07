package http

import (
	"fmt"
	"log"
	"main/internal/adapter/config"
	httpServer "main/internal/adapter/handler/http"
	"main/internal/adapter/storage/mongo"
	"main/internal/adapter/storage/mongo/repository"
	"main/internal/core/service"
)

func HttpMain(
	config *config.Container,
	resource *mongo.Resource,
) {
	// ## dependency injection ##
	// user
	userRepo := repository.NewUserRepository(resource.DB)

	// room
	roomRepo := repository.NewRoomRepository(resource.DB)

	// chat
	chatRepo := repository.NewChatRepository(resource.DB)

	// ## service ##
	// user
	userService := service.NewUserService(
		userRepo,
	)

	// room
	roomService := service.NewRoomService(
		roomRepo,
	)

	// chat
	chatService := service.NewChatService(
		chatRepo,
	)

	// ## handler ##
	userHandler := httpServer.NewUserHandler(userService)

	roomHandler := httpServer.NewRoomHandler(roomService)

	chatHandler := httpServer.NewChatHandler(chatService)

	// router
	fmt.Println("initializing router")
	router, err := httpServer.NewRouter(
		config.HTTP,
		*chatHandler,
		*roomHandler,
		*userHandler,
	)
	if err != nil {
		log.Fatalf("Error initializing router", err)
	}

	// server
	fmt.Printf("Server listening on port %s\n", config.HTTP.Port)
	listenAddr := fmt.Sprintf(":%s", config.HTTP.Port)

	err = router.Serve(listenAddr)
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
