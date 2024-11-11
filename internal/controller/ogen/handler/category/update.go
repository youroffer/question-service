package category

import (
	"context"
	"errors"

	api "github.com/himmel520/question-service/api/oas"
	"github.com/himmel520/question-service/internal/entity"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminCategoriesIDPut(ctx context.Context, req *api.CategoryPut, params api.V1AdminCategoriesIDPutParams) (api.V1AdminCategoriesIDPutRes, error){
	newCategory := &entity.CategoryUpdate{
		Title:  entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().Set},
		Public: entity.Optional[bool]{Value: req.GetPublic().Value, Set: req.GetPublic().Set},
	}

	if !newCategory.IsSet() {
		return &api.V1AdminCategoriesIDPutBadRequest{Message: "no changes"}, nil
	}

	category, err := h.uc.Update(ctx, params.ID, newCategory)
	switch {
	case errors.Is(err, repoerr.ErrCategoryNotFound):
		return &api.V1AdminCategoriesIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrCategoryExists):
		return &api.V1AdminCategoriesIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}

	return entity.CategoryToApi(category), nil
}
