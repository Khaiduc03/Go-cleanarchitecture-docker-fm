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
    NameFeedBack string     `gorm:"column:name_feed_back;type:varchar(100);default:''"`
    Room         int64      `gorm:"column:room;type:bigint;default:0"`
    Description  string     `gorm:"column:description;type:text;default:''"`
    Status       string     `gorm:"column:status;type:varchar(24);default:'SEND'"`
    TimeReq      time.Time  `gorm:"column:time_req"`
    TimeStarted  int64      `gorm:"column:time_started"`
    TimeFinish   int64      `gorm:"column:time_finish"`
    Category     []Category `gorm:"many2many:feedback_categories;"`
}

func (FeedBack) TableName() string {
	return "FEEDBACK"
}
