package db

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis redis.UniversalClient

func InitRedis() {
	url := os.Getenv("REDIS_URL")
	isCluster := os.Getenv("REDIS_CLUSTER")
	var ctx context.Context = context.Background()

	if Redis != nil {
		return
	}
	if isCluster == "true" {
		Redis = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{url},
		})
	} else {
		opts, err := redis.ParseURL(url)
		if err != nil {
			log.Fatal("Error parsing Redis: ", err)
		}
		Redis = redis.NewClient(opts)
	}

	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		log.Printf("Error connecting to Redis: %v", err)
	} else {
		log.Println("Redis connection established")
	}
}
