package entities

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	CategoryName string   `gorm:"column:category_name;type:varchar(100);default:''"`
	FeedbackID   uint     `gorm:"column:feedback_id"`
	FeedBack     FeedBack `gorm:"foreignKey:FeedbackID;references:ID"`
	Rating       uint     `gorm:"column:rating;type:int;default:0"`
	Description  string   `gorm:"column:description;type:text;default:''"`
	UserID       uint     `gorm:"column:user_id"`
	User         User     `gorm:"foreignKey:UserID;references:ID"`
}

func (Rating) TableName() string {
	return "RATING"
}
