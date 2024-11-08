package category

import (
	"context"
	"errors"


	api "github.com/himmel520/question-service/api/oas"
	"github.com/himmel520/question-service/internal/entity"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1CategoriesPost(ctx context.Context, req *api.CategoryInput) (api.V1CategoriesPostRes, error) {
	newCategory, err := h.uc.Create(ctx, &entity.Category{
		Title: req.GetTitle(),
	})

	// TODO: добавить обработку ошибок
	switch{
	case errors.Is(err, repoerr.ErrCategoryExists):
		return &api.V1CategoriesPostConflict{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return newCategory.CategoryToApi(), nil
}
