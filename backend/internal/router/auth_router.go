package router

import (
	"github.com/ahmadzakyarifin/echo-pos-admin/internal/handler"
	"github.com/ahmadzakyarifin/echo-pos-admin/internal/repository"
	"github.com/ahmadzakyarifin/echo-pos-admin/internal/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func SetupAuthRouter(db *sqlx.DB, server *echo.Echo) {
	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	server.POST("/login", authHandler.Login)
}
