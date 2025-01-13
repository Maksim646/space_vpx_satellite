package postgresql

import (
	"context"
	"database/sql"
	"errors"

	//"strings"
	//"time"

	"github.com/Maksim646/space_vpx_satellite/internal/database/postgresql"
	"github.com/Maksim646/space_vpx_satellite/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type UserRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserRepository {
	return &UserRepository{sqalxConn: sqalxConn}
}

func (r *UserRepository) GetUserBuilder() sq.SelectBuilder {
	builder := postgresql.Builder.Select(
		"users.id",
		"users.password_hash",
		"users.name",
		"users.email",
		"users.created_at",
		"users.updated_at",
	).
		From("users")

	return builder
}

func (r *UserRepository) CreateUser(ctx context.Context, email string, name string, password string) (string, error) {
	query, params, err := postgresql.Builder.Insert("users").
		Columns(
			"email",
			"name",
			"password_hash",
		).
		Values(
			email,
			name,
			password,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", err
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newUserID string
	err = r.sqalxConn.GetContext(ctx, &newUserID, query, params...)
	return newUserID, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	builder := r.GetUserBuilder()
	builder = builder.Where(sq.Eq{"users.id": userID})
	return r.GetUser(ctx, builder)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	builder := r.GetUserBuilder()
	builder = builder.Where(sq.Eq{"users.email": email})
	return r.GetUser(ctx, builder)
}

func (r *UserRepository) GetUser(ctx context.Context, builder sq.SelectBuilder) (model.User, error) {
	var user model.User
	query, params, err := builder.ToSql()
	if err != nil {
		return user, err
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &user, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, model.ErrUserNotFound
		}
	}

	return user, nil
}
