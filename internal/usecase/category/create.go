package category

import (
	"context"

	"github.com/himmel520/question-service/internal/entity"
)

func(uc *CategoryUC) Create(ctx context.Context, category *entity.Category) (*entity.Category, error){
	return uc.repo.Create(ctx, uc.db.DB(), category)
}