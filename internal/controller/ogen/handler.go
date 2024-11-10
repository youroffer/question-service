package ogen

import (
	"context"

	api "github.com/himmel520/question-service/api/oas"
)

type (
	Handler struct {
		Auth
		Error
		Category
	}

	Auth interface {
		HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error)
	}

	Error interface {
		NewError(ctx context.Context, err error) *api.ErrorStatusCode
	}

	Category interface {
		V1AdminCategoriesGet(ctx context.Context, params api.V1AdminCategoriesGetParams) (api.V1AdminCategoriesGetRes, error)
		V1AdminCategoriesPost(ctx context.Context, req *api.CategoryPost) (api.V1AdminCategoriesPostRes, error)
		V1AdminCategoriesCategoryIDDelete(ctx context.Context, params api.V1AdminCategoriesCategoryIDDeleteParams) (api.V1AdminCategoriesCategoryIDDeleteRes, error)
		V1AdminCategoriesCategoryIDPut(ctx context.Context, req *api.CategoryPut, params api.V1AdminCategoriesCategoryIDPutParams) (api.V1AdminCategoriesCategoryIDPutRes, error)
	}
)

type HandlerParams struct {
	Auth
	Error
	Category
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		Auth: params.Auth,
		Error: params.Error,
		Category: params.Category,
	}
}
