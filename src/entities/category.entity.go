package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"column:category_name; type:varchar(100); default:''"`
}

func (Category) TableName() string {
	return "CATEGORY"
}
