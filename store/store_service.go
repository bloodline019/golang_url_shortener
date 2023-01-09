package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Структура для хранения Redis клиента
type StorageService struct {
	redisClient *redis.Client
}

// Переменные для хранения контекста и экземпляра структуры
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// Время жизни короткого URL в Redis
const CacheDuration = 24 * time.Hour

// Функция для инициализации Redis клиента
func InitializeStore() *StorageService {
	// Базовая конфигурация Redis клиента
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Проверка соединения с Redis
	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Error init or connecting to redis: %s", err))
	}

	fmt.Printf("Redis init and started successfully with message: %s\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

// Функция для сохранения соответствия короткого и полного URL в Redis (key/value хранилище)
func SaveUrlMapping(shortUrl string, longUrl string) {
	err := storeService.redisClient.Set(ctx, shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Writing in Redis caused an error: %s", err))
	}
}

// Функция для получения полного URL по короткому из Redis
func RetrieveInitialUrl(shortUrl string) string {
	longUrl, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("This key is not exist: %s", shortUrl))
	}
	return longUrl
}
