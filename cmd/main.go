package main

import (
	"tariff-calculator/tariff-api/pkg/config"
	"tariff-calculator/tariff-api/pkg/handler"
	"tariff-calculator/tariff-api/pkg/repository"
	"tariff-calculator/tariff-api/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация репозитория, сервиса и хендлера
	repo := repository.NewAuthRepository(cfg)
	authService := service.NewAuthService(repo)
	authHandler := handler.NewAuthHandler(authService)

	// Инициализация Gin
	r := gin.Default()

	// Маршруты
	r.POST("/login", authHandler.Login)
	r.POST("/refresh-token", authHandler.RefreshToken)
	r.POST("/logout", authHandler.Logout)

	// Запуск сервера
	r.Run(":8080")
}
