package services

import (
	"context"
	"encoding/json"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"metaverse/internal/core/domain"
	"metaverse/internal/core/ports"
	"metaverse/internal/core/services/rabbitmq"
	"time"
)

var (
	publishingErr = errors.New("error tiying to publishing msg")
	InvalidData   = errors.New("invalid data")
)

type UserService struct {
	commandRepository ports.UserCommandRepository
	queryRepository   ports.UserQueryRepository
	rabbitMQ          rabbitmq.RabbitClient
}

func NewUserService(commandRepository ports.UserCommandRepository, queryRepository ports.UserQueryRepository, rabbitMQ rabbitmq.RabbitClient) *UserService {
	return &UserService{
		commandRepository: commandRepository,
		queryRepository:   queryRepository,
		rabbitMQ:          rabbitMQ,
	}
}

func (u *UserService) GetAll() error {
	return nil
}

func (u *UserService) Login(email, password string) error {
	ctx := context.Background()
	if email == "" || password == "" {
		return InvalidData
	}
	err := u.queryRepository.Login(ctx, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Register(name, dni, username, email, password string, age int8) error {
	ctx := context.Background()

	newUser, err := domain.NewUser(name, dni, username, email, password, age)
	if err != nil {
		return err
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {

		return err
	}
	log.Println(userJSON)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := u.rabbitMQ.Send(ctx, "users_events", "users.created.ar", amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         userJSON,
	}); err != nil {
		panic(err)
	}

	log.Println(u.rabbitMQ)
	return u.commandRepository.Register(ctx, newUser.GetID(), name, dni, username, email, password, age)
}
