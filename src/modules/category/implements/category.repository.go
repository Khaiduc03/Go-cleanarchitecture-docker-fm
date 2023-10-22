package categoryImpl

import (
	"FM/src/entities"
	"FM/src/modules/category"
	"FM/src/modules/category/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	*gorm.DB
}

func NewCategoryRepositoryImpl(DB *gorm.DB) category.CategoryRepository {
	return &CategoryRepositoryImpl{DB: DB}
}

func (categoryRepository *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]entities.Category, error) {
	var categories []entities.Category
	err := categoryRepository.DB.Find(&categories).Error
	return categories, err
}

func (categoryRepository *CategoryRepositoryImpl) FindById(ctx context.Context, id int) (entities.Category, error) {
	var category entities.Category
	isExist := categoryRepository.DB.WithContext(ctx).Where("id = ?", id).Find(&category)
	if isExist.RowsAffected == 0 {
		return entities.Category{}, errors.New("category not found")
	}
	return category, nil
}

func (categoryRepository *CategoryRepositoryImpl) Create(ctx context.Context, name string) (string, error) {
	var category entities.Category
	isExist := categoryRepository.DB.WithContext(ctx).Where("category_name = ? ", name).Find(&category)
	if isExist.RowsAffected != 0 {
		return "", errors.New("category is exist")
	}

	category = entities.Category{CategoryName: name}
	err := categoryRepository.DB.WithContext(ctx).Create(&category).Error
	if err != nil {
		return "", err
	}

	return "Create category successfully", nil
}

func (categoryRepository *CategoryRepositoryImpl) Update(ctx context.Context, model modelCategory.UpdateCategoryReq) (string, error) {
	var category entities.Category

	result := categoryRepository.DB.WithContext(ctx).Where("id = ?", model.ID).First(&category)
	if result.RowsAffected == 0 {
		return "", errors.New("category not found")
	}

	if err := categoryRepository.DB.WithContext(ctx).Save(&category).Error; err != nil {
		return "", err
	}
	return "Update category successfull", nil
}

func (categoryRepository *CategoryRepositoryImpl) Delete(ctx context.Context, id int) (string, error) {
	var category entities.Category
	result := categoryRepository.DB.WithContext(ctx).Where("id = ?", id).First(&category)
	if result.RowsAffected == 0 {
		return "", errors.New("category not found")
	}

	if err := categoryRepository.DB.WithContext(ctx).Delete(&category).Error; err != nil {
		return "", err
	}
	return "Delete category successfull", nil
}
