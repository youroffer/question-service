package category

import (
	"context"

	"github.com/himmel520/question-service/internal/entity"
)

func(c *CategoryUC) Create(ctx context.Context, category *entity.Category) (*entity.Category, error){
	return c.repo.Create(ctx, c.db.DB(), category)
}