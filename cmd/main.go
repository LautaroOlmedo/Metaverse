package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"metaverse/internal/core/services"
	"metaverse/internal/core/services/rabbitmq"
	"metaverse/internal/infrastructure/database"
	"metaverse/internal/infrastructure/database/mariadb/command"
	"metaverse/internal/infrastructure/database/mariadb/query"
	"metaverse/internal/infrastructure/http"
	"metaverse/internal/infrastructure/http/handlers"
	"metaverse/settings"
)

func main() {
	myConfig, err := settings.New()
	if err != nil {
		panic(err)
	}
	connection, err := databaseConnection(myConfig)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	conn, err := rabbitmq.ConnectRabbitMQ("lautaro", "secret", "localhost:5672", "users") // myConfig.RabbitMQ.Username, myConfig.RabbitMQ.Password, myConfig.RabbitMQ.Host, myConfig.RabbitMQ.Vhost
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := rabbitmq.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	userCommandRepository := command.NewUserCommandMariaDBRepository(connection)
	userQueryRepository := query.NewUserQueryMariaDBRepository(connection)

	userService := services.NewUserService(userCommandRepository, userQueryRepository, client)

	userHandler := handlers.NewUserHandler(*userService)

	httpServer := http.NewServer(client, *userHandler, *myConfig)
	httpServer.StartServer()
}

func databaseConnection(s *settings.Settings) (*sqlx.DB, error) {
	myContext := context.Background()
	conn, err := database.New(myContext, s)
	if err != nil {
		log.Panicf("failed to start database %s", err)

	}
	return conn, nil
}
