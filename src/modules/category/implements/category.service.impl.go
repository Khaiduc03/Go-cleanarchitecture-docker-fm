package categoryImpl

import (
	"FM/src/entities"
	"FM/src/modules/category"
	modelCategory "FM/src/modules/category/model"
	"context"
)

type CategoryServiceImpl struct {
	category.CategoryRepository
}

func NewCategoryServiceImpl(categoryRepository *category.CategoryRepository) category.CategoryService {
	return &CategoryServiceImpl{CategoryRepository: *categoryRepository}
}

// find all category
func (categoryService *CategoryServiceImpl) FindAll(ctx context.Context) ([]entities.Category, error) {
	return categoryService.CategoryRepository.FindAll(ctx)
}

//find category by id

func (categoryService *CategoryServiceImpl) FindById(ctx context.Context, id int) (entities.Category, error) {
	return categoryService.CategoryRepository.FindById(ctx, id)
}

//create category

func (categoryService *CategoryServiceImpl) Create(ctx context.Context, req modelCategory.CreateCategoryReq) (string, error) {
	response, err := categoryService.CategoryRepository.Create(ctx, req)
	return response, err
}

// update category
func (categoryService *CategoryServiceImpl) Update(ctx context.Context, model modelCategory.UpdateCategoryReq) (string, error) {
	return categoryService.CategoryRepository.Update(ctx, model)
}

//delete category

func (categoryService *CategoryServiceImpl) Delete(ctx context.Context, id int) (string, error) {
	return categoryService.CategoryRepository.Delete(ctx, id)
}

func (categoryService *CategoryServiceImpl) FindAllCategoryByType(ctx context.Context, categoryType string) ([]entities.Category, error) {
	return categoryService.CategoryRepository.FindAllCategoryByType(ctx, categoryType)
}

