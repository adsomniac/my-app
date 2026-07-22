package services

import (
	"context"
	"errors"

	"github.com/adsomniac/my-app/models"
	"github.com/adsomniac/my-app/repositories"
	"github.com/adsomniac/my-app/utils"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("Email atau password salah")
	ErrUserInactive       = errors.New("Akun pengguna tidak aktif")
)

type AuthService struct {
	UserRepository *repositories.UserRepository
	JWTService     *utils.JWTService
}

func NewAuthService(userRepo *repositories.UserRepository, jwtSvc *utils.JWTService) *AuthService {
	return &AuthService{
		UserRepository: userRepo,
		JWTService:     jwtSvc,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.LoginResult, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email dan password wajib diisi")
	}

	user, err := s.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if !user.IsActive {
		return nil, ErrUserInactive
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := s.JWTService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.LoginResult{
		Token: token,
		User:  *user,
	}, nil
}
