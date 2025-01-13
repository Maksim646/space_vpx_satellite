package postgresql

import (
	"context"
	//"database/sql"
	//"errors"
	//"strings"
	//"time"

	"github.com/Maksim646/space_vpx_satellite/internal/database/postgresql"
	"github.com/Maksim646/space_vpx_satellite/internal/model"

	//sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type UserRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserRepository {
	return &UserRepository{sqalxConn: sqalxConn}
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
