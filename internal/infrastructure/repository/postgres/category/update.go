package category

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/question-service/internal/entity"
	"github.com/himmel520/question-service/internal/infrastructure/repository"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *CategoryRepo) Update(ctx context.Context, qe repository.Querier, id int, category *entity.CategoryUpdate) (*entity.Category, error) {
	builder := squirrel.Update("categories").
		Where(squirrel.Eq{"id": id}).
		Suffix(`
		returning 
			id, 
			title, 
			public`).
		PlaceholderFormat(squirrel.Dollar)

	if category.Title.Set {
		builder = builder.Set("title", category.Title.Value)
	}

	if category.Public.Set {
		builder = builder.Set("public", category.Public.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	newCategory := &entity.Category{}
	err = qe.QueryRow(ctx, query, args...).Scan(
		&newCategory.ID, 
		&newCategory.Title, 
		&newCategory.Public)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrCategoryNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.UniqueConstraint {
			return nil, repoerr.ErrCategoryExists
		}
	}

	return newCategory, err
}
