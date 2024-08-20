package handlers

import (
	"github.com/go-redis/redis/v8"
	kafka "github.com/mirjalilova/auth-service-blacklist/pkg/kafka/producer"
	"github.com/mirjalilova/auth-service-blacklist/service"
)

type Handlers struct {
	Auth     *service.AuthService
	User     *service.UserService
	RDB      *redis.Client
	Producer kafka.KafkaProducer
}

func NewHandler(auth *service.AuthService, user *service.UserService, rdb *redis.Client, pr *kafka.KafkaProducer) *Handlers {
	return &Handlers{
		Auth:     auth,
		User:     user,
		RDB:      rdb,
		Producer: *pr,
	}
}
