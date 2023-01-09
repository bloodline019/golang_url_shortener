package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	longLink := "https://www.google.com"
	shortURL := "Jsz4k57oAX"

	// Сохранение соответствия пары ссылок
	SaveUrlMapping(shortURL, longLink)

	// Получение исходной ссылки
	retrievedUrl := RetrieveInitialUrl(shortURL)

	// Проверка на соответствие полученной и исходной ссылок
	assert.Equal(t, longLink, retrievedUrl)
}
