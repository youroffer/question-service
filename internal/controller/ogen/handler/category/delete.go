package category

import (
	"context"

	api "github.com/himmel520/question-service/api/oas"
)

func (h *Handler) V1AdminCategoriesIDDelete(ctx context.Context, params api.V1AdminCategoriesIDDeleteParams) (api.V1AdminCategoriesIDDeleteRes, error) {
	return &api.V1AdminCategoriesIDDeleteNoContent{}, nil
}
