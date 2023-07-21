package postgres

import (
	"context"
	"fmt"
	"go-learn/internal/domain"
	"go-learn/internal/repository/postgres/dao"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r Repository) CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error) {

	var userID uuid.UUID
	usr := dao.FromUserDomain(user)

	err := pgx.BeginTxFunc(ctx, r.conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		id, txErr := r.createUserTx(ctx, tx, usr)
		userID = id
		return txErr
	})

	return userID, err
}

func (r Repository) createUserTx(ctx context.Context, tx pgx.Tx, user dao.User) (uuid.UUID, error) {
	fmt.Println("2-user", user)
	var id uuid.UUID

	query, args, err := r.genSql.
		Insert(dao.UserTableName).
		Columns(dao.ColumnForCreateUser...).
		Values(user.Name).
		ToSql()
	if err != nil {
		return id, err
	}

	err = pgxscan.Get(ctx, tx, &id, query, args...)

	if err != nil {
		return id, err
	}

	return id, nil
}
