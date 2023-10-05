package entities

import (
	"time"

	"gorm.io/gorm"
)

const (
	SEND   = "SEND"
	ACCEPT = "ACCEPT"
	FINISH = "FINISH"
)

type FeedBack struct {
	gorm.Model
	NameFeedBack string     `gorm:"column:NameFeedBack; type:varchar(100); default:''"`
	Room         int64      `gorm:"column:dob; type:varchar(24); default:''"`
	Description  string     `gorm:"column:Description; type:text; default:''"`
	Status       string     `gorm:"column:Status; type:varchar(24); default:'SEND'"`
	TimeReq      time.Time  `gorm:"column:TimeReq;  type:bigint;default:0 "`
	TimeStarted  int64      `gorm:"column:TimeStarted;  type:bigint;  default:0"`
	TimeFinish   int64      `gorm:"column:TimeFinish;  type:bigint;  default:0"`
	Category     []Category `gorm:"many2many:feedback_category;"`
}

func (FeedBack) TableName() string {
	return "FEEDBACK"
}
