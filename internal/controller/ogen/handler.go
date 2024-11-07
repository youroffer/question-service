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
		V1CategoriesPost(ctx context.Context, req *api.CategoryInput) (api.V1CategoriesPostRes, error)
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
