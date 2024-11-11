package category

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/question-service/internal/infrastructure/repository"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *CategoryRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	query, args, err := squirrel.Delete("categories").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	cmdTag, err := qe.Exec(ctx, query, args...)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.FKViolation {
			return repoerr.ErrCategoryDependency
		}
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrCategoryNotFound
	}

	return err
}
