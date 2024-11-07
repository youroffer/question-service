package auth

import (
	"context"

	api "github.com/himmel520/question-service/api/oas"
)

func (h *Handler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	// TODO: реализовать проверку токена
	return ctx, nil
}
