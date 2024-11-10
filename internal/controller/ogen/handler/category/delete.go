package category

import (
	"context"

	api "github.com/himmel520/question-service/api/oas"
)

func (h *Handler) V1AdminCategoriesCategoryIDDelete(ctx context.Context, params api.V1AdminCategoriesCategoryIDDeleteParams) (api.V1AdminCategoriesCategoryIDDeleteRes, error) {
	return &api.V1AdminCategoriesCategoryIDDeleteNoContent{}, nil
}
