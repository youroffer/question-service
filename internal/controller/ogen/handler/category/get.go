package category

import (
	"context"

	api "github.com/himmel520/question-service/api/oas"
)

func (h *Handler) V1AdminCategoriesGet(ctx context.Context, params api.V1AdminCategoriesGetParams) (api.V1AdminCategoriesGetRes, error) {
	return &api.CategoriesResp{}, nil
}
