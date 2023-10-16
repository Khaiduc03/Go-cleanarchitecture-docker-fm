package category

import (
	"FM/src/entities"
	"FM/src/modules/category/model"
	"context"
)

type CategoryService interface {
	FindAll(ctx context.Context) ([]entities.Category, error)
	FindById(ctx context.Context, id int) (entities.Category, error)
	Create(ctx context.Context, name string) (string, error)
	Update(ctx context.Context, model model.UpdateCategoryReq) (string, error)
	Delete(ctx context.Context, id int) (string, error)
}
