package lib

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	r := redis.NewClient(&redis.Options{
		Addr:     "lib:6379",
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
	return &RedisClient{
		Client: r,
	}
}

func (r *RedisClient) GetOtp(otp string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.Client.Get(ctx, otp).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New("key not found")
		}
		return "", fmt.Errorf("lib get failed: %w", err)
	}

	return result, nil
}

func (r *RedisClient) SetOtp(phone string, otp string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Client.Set(ctx, otp, phone, 5*time.Minute).Err()
	if err != nil {
		return errors.New("lib set failed: " + err.Error())
	}
	return nil
}
