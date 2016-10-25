package main

import (
    "log"
	"os"

    "github.com/joho/godotenv"
	"github.com/mattmac4241/api-gateway/service"
)

func main() {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }

	redisAddress := os.Getenv("REDIS_PORT")

	service.REDIS, _ = service.InitRedisClient(redisAddress, "")
	defer service.REDIS.Close()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3003"
	}
	server := service.NewServer()
	server.Run(":" + port)
}
