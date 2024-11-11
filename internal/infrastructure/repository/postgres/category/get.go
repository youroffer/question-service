package category

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/question-service/internal/entity"
	"github.com/himmel520/question-service/internal/infrastructure/repository"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
)

func (r *CategoryRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Category, error) {
	query, _, err := squirrel.Select(
		"id",
		"title",
		"public").
	From("categories").
	Limit(params.Limit).
	Offset(params.Offset).
	ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qe.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	categories := []*entity.Category{}
	for rows.Next() {
		category := &entity.Category{}
		if err := rows.Scan(&category.ID, &category.Title, &category.Public); err != nil {
			return nil, err
		}

		categories = append(categories, category)

	}

	if len(categories) == 0 {
		return nil, repoerr.ErrCategoryNotFound
	}

	return categories, err
}

func (r *CategoryRepo) Count(ctx context.Context, qe repository.Querier) (int, error) {
	var count int
	err := qe.QueryRow(ctx, `select COUNT(*) from categories;`).Scan(&count)
	return count, err

}
