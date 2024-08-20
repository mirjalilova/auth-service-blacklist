package app

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/mirjalilova/auth-service-blacklist/api"
	"github.com/mirjalilova/auth-service-blacklist/api/handlers"
	"github.com/mirjalilova/auth-service-blacklist/internal/config"
	kafka "github.com/mirjalilova/auth-service-blacklist/pkg/kafka/consumer"
	prd "github.com/mirjalilova/auth-service-blacklist/pkg/kafka/producer"
	"github.com/mirjalilova/auth-service-blacklist/pkg/storage/postgres"
	"github.com/mirjalilova/auth-service-blacklist/service"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {

	// Postgres Connection
	db, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		slog.Error("can't connect to db: %v", err)
	}
	defer db.Db.Close()
	slog.Info("Connected to Postgres")

	// Redis Connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		slog.Error("Failed to connect to Redis: %v", err)
	}
	slog.Info("Connected to Redis")

	authService := service.NewAuthService(db)
	userService := service.NewUserService(db)

	// Kafka
	brokers := []string{"kafka:9092"}
	cm := kafka.NewKafkaConsumerManager()
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		slog.Error("Failed to create Kafka producer:", err)
		return
	}

	Reader(brokers, cm, authService, userService)

	// HTTP Server
	h := handlers.NewHandler(authService, userService, rdb, &pr)

	router := api.Engine(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.AUTH_PORT); err != nil {
		slog.Error("can't start server: %v", err)
	}

	slog.Info("REST server started on port %s", cfg.AUTH_PORT)
}
