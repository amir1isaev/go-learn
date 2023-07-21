package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"go-learn/pkg/database/postgres"
	"go-learn/pkg/pagination"
)

type Repository struct {
	conn   *postgres.DB
	genSql squirrel.StatementBuilderType
}

func New(conn *postgres.DB) *Repository {
	return &Repository{
		conn:   conn,
		genSql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r Repository) preparePaginatedQuery(
	gen squirrel.StatementBuilderType,
	tableName string,
	pagination pagination.Pagination,
	searchCondition map[string]interface{},
) (
	countQuery string, countParams []any,
	selectQuery string, selectParams []any,
	err error,
) {
	preparedCountQuery := gen.Select("count(*)").From(tableName)
	if searchCondition != nil {
		preparedCountQuery = preparedCountQuery.
			Where(searchCondition)
	}

	countQuery, countParams, err = preparedCountQuery.ToSql()
	if err != nil {
		return
	}

	preparedQuery := gen.Select("*").From(tableName)
	if searchCondition != nil {
		preparedQuery = preparedQuery.
			Where(searchCondition)
	}

	selectQuery, selectParams, err = preparedQuery.
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		ToSql()
	if err != nil {
		return
	}

	return
}

func (r Repository) addEntitiesRelation(
	ctx context.Context, tx pgx.Tx,
	tableName string,
	data map[string]interface{},
) error {
	query, args, err := r.genSql.
		Insert(tableName).
		SetMap(data).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) removeEntitiesRelation(
	ctx context.Context, tx pgx.Tx,
	tableName string,
	data map[string]interface{},
) error {
	query, args, err := r.genSql.
		Delete(tableName).
		Where(data).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) scanGetQuery(
	ctx context.Context, tx pgx.Tx,
	data interface{}, withCondition bool,
	query string, params ...interface{},
) (err error) {
	if withCondition {
		err = pgxscan.Get(ctx, tx, data, query, params...)
	} else {
		err = pgxscan.Get(ctx, tx, data, query)
	}

	return
}

func (r Repository) scanSelectQuery(
	ctx context.Context, tx pgx.Tx,
	data interface{}, withCondition bool,
	query string, params ...interface{},
) (err error) {
	if withCondition {
		err = pgxscan.Select(ctx, tx, data, query, params...)
	} else {
		err = pgxscan.Select(ctx, tx, data, query)
	}

	return
}
