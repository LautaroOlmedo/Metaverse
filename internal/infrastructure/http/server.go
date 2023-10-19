package http

import (
	"log"
	"metaverse/internal/core/services/rabbitmq"
	"metaverse/internal/infrastructure/http/handlers"
	"metaverse/settings"
	"net/http"
)

type Server struct {
	rabbitMQ    rabbitmq.RabbitClient
	userHandler *handlers.UserHandler
	settings    *settings.Settings
}

func NewServer(rabbitMQ rabbitmq.RabbitClient, userHandler handlers.UserHandler, settings settings.Settings) *Server {
	return &Server{
		rabbitMQ:    rabbitMQ,
		userHandler: &userHandler,
		settings:    &settings,
	}
}

func (s *Server) StartServer() {
	mux := http.NewServeMux()
	err := http.ListenAndServe(s.settings.Port, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
	log.Printf("Server listening on port: %s", s.settings.Port)
}
