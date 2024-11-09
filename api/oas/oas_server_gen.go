// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// V1AdminCategoriesDelete implements DELETE /v1/admin/categories operation.
	//
	// Удаляет категорию по ее уникальному идентификатору.
	//
	// DELETE /v1/admin/categories
	V1AdminCategoriesDelete(ctx context.Context, params V1AdminCategoriesDeleteParams) (V1AdminCategoriesDeleteRes, error)
	// V1AdminCategoriesGet implements GET /v1/admin/categories operation.
	//
	// Возвращает список всех категорий с возможностью
	// пагинации.
	//
	// GET /v1/admin/categories
	V1AdminCategoriesGet(ctx context.Context, params V1AdminCategoriesGetParams) (V1AdminCategoriesGetRes, error)
	// V1AdminCategoriesPost implements POST /v1/admin/categories operation.
	//
	// Создает новую категорию.
	//
	// POST /v1/admin/categories
	V1AdminCategoriesPost(ctx context.Context, req *CategoryInput) (V1AdminCategoriesPostRes, error)
	// V1AdminCategoriesPut implements PUT /v1/admin/categories operation.
	//
	// Обновляет категорию.
	//
	// PUT /v1/admin/categories
	V1AdminCategoriesPut(ctx context.Context, req *CategoryInput) (V1AdminCategoriesPutRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
