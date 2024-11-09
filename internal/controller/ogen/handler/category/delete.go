package category

import (
	"context"

	api "github.com/himmel520/question-service/api/oas"
)

func (h *Handler) V1AdminCategoriesDelete(ctx context.Context, params api.V1AdminCategoriesDeleteParams) (api.V1AdminCategoriesDeleteRes, error) {
	return &api.V1AdminCategoriesDeleteNoContent{}, nil
}
