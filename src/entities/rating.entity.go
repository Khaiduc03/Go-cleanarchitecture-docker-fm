package entities

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	CategoryName string   `gorm:"column:category_name; type:varchar(100); default:''"`
	IDFeedback   FeedBack `gorm:"foreignKey:IDFeedback;references:ID;"`
	Rating       uint     `gorm:"column:rating; type:int; default:0"`
	Description  string   `gorm:"column:description; type:text; default:''"`
}

func (Rating) TableName() string {
	return "RATING"
}
