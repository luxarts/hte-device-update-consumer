package main

import (
	"context"
	"hte-device-update-consumer/internal/controller"
	"hte-device-update-consumer/internal/defines"
	"hte-device-update-consumer/internal/metrics"
	"hte-device-update-consumer/internal/repository"
	"hte-device-update-consumer/internal/service"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/go-resty/resty/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	metrics.StartServer()

	// Redis Client
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(defines.EnvRedisHost),
		Password: os.Getenv(defines.EnvRedisPassword),
	})
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Error ping Redis: %+v\n", err)
	}

	// Rest Client
	restClient := resty.New()

	// Dependency injection
	locationRepo := repository.NewLocationRepository(restClient)
	statusRepo := repository.NewStatusRepository(restClient)
	svc := service.NewMessageService(locationRepo, statusRepo)
	ctrl := controller.NewMessageController(svc)

	// Read queue
	log.Printf("Polling queue %s\n", defines.QueueDeviceUpdate)

	for {
		msg, err := redisClient.LPop(ctx, defines.QueueDeviceUpdate).Result()
		if err != nil {
			if err.Error() == "redis: nil" {
				continue
			}
			log.Printf("Error receiving msg: %+v\n", err)
		}

		ctrl.Process(&msg)
	}
}
