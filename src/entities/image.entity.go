package entities

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	PublicID    string   `gorm:"column:public_id; type:varchar(100); default:''"`
	Url         string   `gorm:"column:url; type:varchar(100); default:''"`
	IDFeedBacks FeedBack `gorm:"foreignKey:IDFeedBacks;references:ID;"`
}

func (Image) TableName() string {
	return "IMAGE"
}
