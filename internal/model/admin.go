package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrAdminNotFound = errors.New("admin not found")
	AdminNotFound    = "admin not found"
)

type Admin struct {
	ID string `db:"id"`

	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`

	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

type IAdminRepository interface {
	CreateAdmin(ctx context.Context, email string, name string, password string) (string, error)

	GetAdminByEmail(ctx context.Context, email string) (Admin, error)
	GetAdminByID(ctx context.Context, userID string) (Admin, error)
}

type IAdminUsecase interface {
	CreateAdmin(ctx context.Context, email string, name string, password string) (string, error)

	GetAdminByEmail(ctx context.Context, email string) (Admin, error)
	GetAdminByID(ctx context.Context, userID string) (Admin, error)
}
