package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	r := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	return &RedisClient{
		client: r,
	}
}

func (r *RedisClient) GetHandler(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", fmt.Errorf("result not found")
		}
		return "", fmt.Errorf("error getting value: %v", err)
	}

	return url, nil
}

func (r *RedisClient) SetHandler(url string) (string, error) {
	shortKey := generateShortKey(4)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url, err := r.client.Set(ctx, shortKey, url, 24*time.Hour).Result()

	if err != nil {
		return "", fmt.Errorf("URL already exists or error setting value: %v", err)
	}

	return shortKey, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortKey(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
