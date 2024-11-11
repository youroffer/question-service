package category

import (
	"context"
	"errors"

	api "github.com/himmel520/question-service/api/oas"
	"github.com/himmel520/question-service/internal/controller/ogen"
	"github.com/himmel520/question-service/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/question-service/internal/usecase"
)

func (h *Handler) V1AdminCategoriesGet(ctx context.Context, params api.V1AdminCategoriesGetParams) (api.V1AdminCategoriesGetRes, error) {
	categoriesResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})
	
	switch {
	case errors.Is(err, repoerr.ErrCategoryNotFound):
		return &api.V1AdminCategoriesGetNotFound{Message: err.Error()}, nil
	case err != nil:
		h.log.Error(err)
		return nil, err
	}

	return categoriesResp.ToApi(), nil
}
