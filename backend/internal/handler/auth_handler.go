package handler

import (
	"github.com/ahmadzakyarifin/echo-pos-admin/internal/dto/auth"
	"github.com/ahmadzakyarifin/echo-pos-admin/internal/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(server echo.Context) error {
	var req auth.LoginRequest
	if err := server.Bind(&req); err != nil {
		return server.JSON(400, auth.LoginErrorResponse{
			Status:  "error",
			Message: "Invalid request",
		})
	}
	token, role, err := h.authUsecase.Login(req.Username, req.Email, req.Password)
	if err != nil {
		return server.JSON(401, auth.LoginErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return server.JSON(200, auth.LoginSuccessResponse{
		Status: "success",
		Token:  token,
		Role:   role,
	})
}
