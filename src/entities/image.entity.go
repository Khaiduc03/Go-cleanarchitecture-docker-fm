package entities

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	//PublicID   string   `gorm:"column:public_id;type:varchar(100);default:''"`
	Url        string   `gorm:"column:url;type:varchar(256);default:''"`
	FeedbackID uint     `gorm:"column:feedback_id"`
	FeedBack   FeedBack `gorm:"foreignKey:FeedbackID;references:ID"`
}

func (Image) TableName() string {
	return "IMAGE"
}
