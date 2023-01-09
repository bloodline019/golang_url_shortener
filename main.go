package main

import (
	"fmt"
	"github.com/bloodline019/golang_url_shortener/routes"
	"github.com/bloodline019/golang_url_shortener/store"
	"github.com/gin-gonic/gin"
)

// Создаем стандартный роутер Gin
var r = gin.Default()

func main() {
	// Задаем папку с шаблонами
	r.LoadHTMLGlob("templates/*.html")
	// Инициализируем хранилище Redis
	store.InitializeStore()
	// Инициализируем маршруты API
	routes.InitializeRoutes(r)
	// Запускаем сервер на порту 8080
	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprint("Error starting server: ", err))
	}
}
