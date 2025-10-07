package main

import (
	"tariff-api/internal/config"
	"tariff-api/internal/handler"
	mw "tariff-api/internal/middleware"
	"tariff-api/internal/repository"
	"tariff-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	repo := repository.NewAuthRepository(cfg)
	authService := service.NewAuthService(repo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()
	r.Use(mw.CORS())

	r.POST("/login", authHandler.Login)
	r.POST("/refresh-token", authHandler.RefreshToken)
	r.POST("/logout", authHandler.Logout)

	r.Run(":8081")
}
