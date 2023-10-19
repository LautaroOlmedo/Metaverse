package services

import (
	"context"
	"encoding/json"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"metaverse/internal/core/domain"
	"metaverse/internal/core/ports"
	"metaverse/internal/core/services/rabbitmq"
	"time"
)

var (
	publishingErr = errors.New("error tiying to publishing msg")
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
	err := u.queryRepository.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Register(name, dni, username, email, password string, age int8) error {
	ctx := context.Background()
	err := u.commandRepository.Register(ctx, name, dni, username, email, password, age)
	if err != nil {
		return err
	}

	newUser, err := domain.NewUser(name, dni, username, email, password, age)

	userJSON, err := json.Marshal(newUser)
	if err != nil {

		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := u.rabbitMQ.Send(ctx, "dollar", "dollar.user.created.*", amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         userJSON,
	}); err != nil {
		return publishingErr
	}
	return nil
}
