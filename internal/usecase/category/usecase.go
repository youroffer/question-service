package category

import (
	"context"

	"github.com/himmel520/question-service/internal/entity"
	"github.com/himmel520/question-service/internal/infrastructure/repository"
)

type (
	CategoryUC struct {
		db   DBXT
		repo CategoryRepo
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	CategoryRepo interface {
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Category, error)
		Create(ctx context.Context, qe repository.Querier, category *entity.Category) (*entity.Category, error)
		Update(ctx context.Context, qe repository.Querier, id int, category *entity.CategoryUpdate) (*entity.Category, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Count(ctx context.Context, qe repository.Querier) (int, error)
	}
)

func New(db DBXT, repo CategoryRepo) *CategoryUC {
	return &CategoryUC{db: db, repo: repo}
}
