package category

import (
	"context"
	"errors"

	api "github.com/himmel520/question-service/api/oas"
	"github.com/himmel520/question-service/internal/entity"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminCategoriesPost(ctx context.Context, req *api.CategoryPost) (api.V1AdminCategoriesPostRes, error) {
	newCategory, err := h.uc.Create(ctx, &entity.Category{
		Title:  req.GetTitle(),
		Public: req.GetPublic(),
	})

	// TODO: добавить обработку ошибок
	switch {
	case errors.Is(err, repoerr.ErrCategoryExists):
		return &api.V1AdminCategoriesPostConflict{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return entity.CategoryToApi(newCategory), nil
}
