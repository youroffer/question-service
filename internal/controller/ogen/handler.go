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
		V1AdminCategoriesIDDelete(ctx context.Context, params api.V1AdminCategoriesIDDeleteParams) (api.V1AdminCategoriesIDDeleteRes, error)
		V1AdminCategoriesIDPut(ctx context.Context, req *api.CategoryPut, params api.V1AdminCategoriesIDPutParams) (api.V1AdminCategoriesIDPutRes, error)
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
