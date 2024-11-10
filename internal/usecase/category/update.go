package category

import (
	"context"

	"github.com/himmel520/question-service/internal/entity"
)

func (uc *CategoryUC) Update(ctx context.Context, id int, category *entity.CategoryUpdate) (*entity.Category, error) {
	return uc.repo.Update(ctx, uc.db.DB(), id, category)
}
