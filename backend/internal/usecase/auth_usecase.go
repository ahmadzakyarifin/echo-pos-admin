package usecase

import (
	"errors"

	"github.com/ahmadzakyarifin/echo-pos-admin/internal/pkg"
	"github.com/ahmadzakyarifin/echo-pos-admin/internal/repository"
)

type AuthUsecase interface {
	Login(username, email, password string) (string, string, error)
}

type authUsecase struct {
	authRepo repository.AuthRepository
}

func NewAuthUsecase(authRepo repository.AuthRepository) AuthUsecase {
	return &authUsecase{authRepo: authRepo}
}

func (u *authUsecase) Login(username, email, password string) (string, string, error) {
	user, err := u.authRepo.FindUserByIdentifier(username, email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("Email atau Username tidak terdaftar")
	}

	if !pkg.VerifyPassword(password, user.Password) {
		return "", "", errors.New("Password yang Anda masukkan salah")
	}

	token, err := pkg.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	return token, user.Role, nil
}
