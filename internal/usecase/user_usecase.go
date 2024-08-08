package usecase

import (
	"context"
	"time"

	"errors"
	"jwtGolang/internal/domain"
	"jwtGolang/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type UserUsecaseInterface interface {
	Login
	Register
	Welcome
}

type Login interface {
	Login(ctx context.Context, user domain.User) (string, error)
}

type Register interface {
	Register(ctx context.Context, user *domain.User) error
}
type Welcome interface {
	Welcome(ctx context.Context, tokenString string) (string, error)
}

type UserUsecase struct {
	UserRepo repository.UserRepositoryInterface
}

func NewUserusecase(userRepo repository.UserRepositoryInterface) UserUsecase {
	return UserUsecase{
		UserRepo: userRepo,
	}
}

func (u UserUsecase) Register(ctx context.Context, user *domain.User) error {
	_, err := u.UserRepo.CreateUser(ctx, user)
	return err
}

func (u UserUsecase) Login(ctx context.Context, user domain.User) (string, error) {
	err := u.UserRepo.VerifyPassword(ctx, user.Username, user.Password)
	if err != nil {
		return "", err
	}

	expireTime := time.Now().Add(25 * time.Minute)
	claim := &domain.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Subject:   user.Username,
		},
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u UserUsecase) Welcome(ctx context.Context, tokenString string) (string, error) {
	claims := &domain.Claims{
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}
	return "Welcome " + claims.RegisteredClaims.Subject, nil
}
