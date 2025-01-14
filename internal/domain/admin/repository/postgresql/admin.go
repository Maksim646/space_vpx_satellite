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

type AdminRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IAdminRepository {
	return &AdminRepository{sqalxConn: sqalxConn}
}

func (r *AdminRepository) GetAdminBuilder() sq.SelectBuilder {
	builder := postgresql.Builder.Select(
		"admins.id",
		"admins.password_hash",
		"admins.name",
		"admins.email",
		"admins.created_at",
		"admins.updated_at",
	).
		From("admins")

	return builder
}

func (r *AdminRepository) CreateAdmin(ctx context.Context, email string, name string, password string) (string, error) {
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

	var newAdminID string
	err = r.sqalxConn.GetContext(ctx, &newAdminID, query, params...)
	return newAdminID, err
}

func (r *AdminRepository) GetAdminByID(ctx context.Context, userID string) (model.Admin, error) {
	builder := r.GetAdminBuilder()
	builder = builder.Where(sq.Eq{"admins.id": userID})
	return r.GetAdmin(ctx, builder)
}

func (r *AdminRepository) GetAdminByEmail(ctx context.Context, email string) (model.Admin, error) {
	builder := r.GetAdminBuilder()
	builder = builder.Where(sq.Eq{"admins.email": email})
	return r.GetAdmin(ctx, builder)
}

func (r *AdminRepository) GetAdmin(ctx context.Context, builder sq.SelectBuilder) (model.Admin, error) {
	var admin model.Admin
	query, params, err := builder.ToSql()
	if err != nil {
		return admin, err
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &admin, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return admin, model.ErrUserNotFound
		}
	}

	return admin, nil
}
