package connections

import (
	"conn/internal/api/handler"
	interface17 "conn/internal/interface"
	"conn/internal/redis/adjust"
	redis17 "conn/internal/redis/methods"
	"conn/internal/redis/services"
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis17.Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &redis17.Redis{R: client, C: ctx}
}

func NewAdjust() interface17.Origin {
	r := NewRedis()
	return &adjust.Adjust{R: r, Generated: map[int]bool{}, Rng: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func NewService()*services.Services{
	a:=NewAdjust()
	return &services.Services{O: a}
}

func NewHandler()*handler.Handler{
	h:=NewService()
	return &handler.Handler{S: h}
}