package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID string `db:"id"`

	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`

	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

type IUserRepository interface {
	CreateUser(ctx context.Context, email string, name string, password string) (string, error)

	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, email string, name string, password string) (string, error)

	GetUserByEmail(ctx context.Context, email string) (User, error)
}
