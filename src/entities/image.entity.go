package entities

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	PublicID   string   `gorm:"column:public_id;type:varchar(100);default:''"`
	Url        string   `gorm:"column:url;type:varchar(100);default:''"`
	IDFeedback uint     `gorm:"column:id_feedback"`
	FeedBack   FeedBack `gorm:"foreignKey:IDFeedback;references:ID"`
}

func (Image) TableName() string {
	return "IMAGE"
}
