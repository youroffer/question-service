package category

import (
	"context"

	"github.com/himmel520/question-service/internal/entity"
)

type (
	Handler struct{
		uc CategoryUsecase
	}

	CategoryUsecase interface {
		Create(ctx context.Context, category *entity.Category) (*entity.Category, error)
	}
)

func New(uc CategoryUsecase) *Handler {
	return &Handler{uc: uc}
}
