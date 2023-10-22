package category

import (
	"FM/src/entities"
	"FM/src/modules/category/model"
	"context"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]entities.Category, error)
	FindAllCategoryByType(ctx context.Context, categoryType string) ([]entities.Category, error)
	FindById(ctx context.Context, id int) (entities.Category, error)
	Create(ctx context.Context, model modelCategory.CreateCategoryReq) (string, error)
	Update(ctx context.Context,model modelCategory.UpdateCategoryReq) (string, error)
	Delete(ctx context.Context, id int) (string, error)
	
}
