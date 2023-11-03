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
	if err := s.rabbitMQ.CreateQueue("users_created", true, false); err != nil {
		panic(err)
	}

	if err := s.rabbitMQ.CreateBinding("users_created", "users.created.*", "users_events"); err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/users", s.userHandler)

	log.Printf("Server listening on port: %s", s.settings.Port)
	err := http.ListenAndServe(s.settings.Port, mux) // ---> blocking function
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
