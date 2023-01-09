package handler

import (
	"github.com/bloodline019/golang_url_shortener/store"
	"github.com/bloodline019/golang_url_shortener/url_generator"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)

type UrlShortenerRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

// IPv4 ардес сервера в локальной сети для доступа с других устройств
const host = "http://localhost:8080/"

// CreateShortUrl Создание короткой ссылки и соответствия с исходной в хранилище Redis
// Возвращаем короткую ссылку в формате JSON
func CreateShortUrl(c *gin.Context) {
	var requestData UrlShortenerRequest
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := url_generator.GenerateShortUrl(requestData.LongUrl, requestData.UserId)
	store.SaveUrlMapping(shortUrl, requestData.LongUrl)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	longUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(http.StatusFound, longUrl)
}

// CreateQrCodeLink Функция обработки запроса на создание QR-кода
func CreateQrCodeLink(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "qrcode.html", nil)
	}
	if c.Request.Method == "POST" {
		longUrl := "https://" + c.PostForm("dataString")
		userId := "wqer7c6-7wer-qnof-fneiun"

		shortUrl := url_generator.GenerateShortUrl(longUrl, userId)
		store.SaveUrlMapping(shortUrl, longUrl)

		png, err := qrcode.Encode(host+shortUrl, qrcode.Medium, 256)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "image/png", png)
	}
}
