package main

import (
	"tariff-calculator/tariff-api/pkg/config"
	"tariff-calculator/tariff-api/pkg/handler"
	"tariff-calculator/tariff-api/pkg/repository"
	"tariff-calculator/tariff-api/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	repo := repository.NewAuthRepository(cfg)
	authService := service.NewAuthService(repo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/login", authHandler.Login)
	r.POST("/refresh-token", authHandler.RefreshToken)
	r.POST("/logout", authHandler.Logout)

	r.Run(":8081")
}
