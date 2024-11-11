package category

import (
	"context"
	"errors"

	api "github.com/himmel520/question-service/api/oas"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminCategoriesIDDelete(ctx context.Context, params api.V1AdminCategoriesIDDeleteParams) (api.V1AdminCategoriesIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.ID)
	switch {
	case errors.Is(err, repoerr.ErrCategoryNotFound):
		return &api.V1AdminCategoriesIDDeleteNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrCategoryDependency):
		return &api.V1AdminCategoriesIDDeleteConflict{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return &api.V1AdminCategoriesIDDeleteNoContent{}, nil
}
