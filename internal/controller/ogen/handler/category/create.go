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

	if err != nil {
		if errors.Is(err, repoerr.ErrCategoryExists) {
			return &api.V1CategoriesPostBadRequest{
				Message: err.Error()}, nil
		}
		return nil, err
	}

	return newCategory.CategoryToApi(), err
}
