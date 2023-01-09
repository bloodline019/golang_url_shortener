package routes

import (
	"github.com/bloodline019/golang_url_shortener/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Задаем маршруты и обработчики
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This server is the URL Shortener API",
		})
	})

	r.POST("/create_short_url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	r.Any("/qrcode", func(c *gin.Context) {
		handler.CreateQrCodeLink(c)
	})
}
