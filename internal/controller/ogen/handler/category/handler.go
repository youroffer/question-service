package category

import (
	"context"

	"github.com/himmel520/question-service/internal/entity"
	"github.com/sirupsen/logrus"
)

type (
	Handler struct {
		uc  CategoryUsecase
		log *logrus.Logger
	}

	CategoryUsecase interface {
		Create(ctx context.Context, category *entity.Category) (*entity.Category, error)
		Update(ctx context.Context, id int, category *entity.CategoryUpdate) (*entity.Category, error)
	}
)

func New(uc CategoryUsecase, log *logrus.Logger) *Handler {
	return &Handler{uc: uc, log: log}
}
