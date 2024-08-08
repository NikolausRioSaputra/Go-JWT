package repository

import (
	"context"
	"database/sql"
	"errors"
	"jwtGolang/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	CreateUser
	GetUser
	VerifyPassword
}

type CreateUser interface {
	CreateUser(ctx context.Context, user *domain.User) (domain.User, error)
}

type GetUser interface {
	GetUser(ctx context.Context, username string) (domain.User, error)
}

type VerifyPassword interface {
	VerifyPassword(ctx context.Context, username, password string) error
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser implements UserRepositoryInterface.
func (repo *UserRepository) CreateUser(ctx context.Context, user *domain.User) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	query := `INSERT INTO users(username, password) VALUES($1, $2) RETURNING id`
	err = repo.DB.QueryRowContext(ctx, query, user.Username, string(hashedPassword)).Scan(&user.ID)
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}

// GetUser implements UserRepositoryInterface.
func (repo *UserRepository) GetUser(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := repo.DB.QueryRowContext(ctx, "SELECT username, password FROM users WHERE username = $1", username).Scan(&user.Username, &user.Password)
	return user, err
}

func (repo *UserRepository) VerifyPassword(ctx context.Context, username, password string) error {
	user, err := repo.GetUser(ctx, username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid username or password")
	}
	return nil
}
