package category

import (
	"context"

	"github.com/himmel520/question-service/internal/entity"
	"github.com/himmel520/question-service/internal/infrastructure/repository"
	"github.com/himmel520/question-service/internal/lib/paging"
	"github.com/himmel520/question-service/internal/usecase"
)

func (uc *CategoryUC) Get(ctx context.Context, params usecase.PageParams) (*entity.CategoriesResp, error) {
	categories, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  params.PerPage,
		Offset: params.Page * params.PerPage})
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, err
	}

	return &entity.CategoriesResp{
		Data:    categories,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}
