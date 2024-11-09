package category

import (
	"context"

	api "github.com/himmel520/question-service/api/oas"
)

func (h *Handler) V1AdminCategoriesPut(ctx context.Context, req *api.CategoryInput) (api.V1AdminCategoriesPutRes, error) {
	return &api.Category{}, nil
}
