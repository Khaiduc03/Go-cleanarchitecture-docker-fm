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
	Description  string     `gorm:"column:description;type:text;default:''"`
	Status       string     `gorm:"column:status;type:varchar(24);default:'SEND'"`
	TimeReq      time.Time  `gorm:"column:time_req"`
	TimeStarted  int64      `gorm:"column:time_started"`
	TimeFinish   int64      `gorm:"column:time_finish"`
	Category     []Category `gorm:"many2many:feedback_category;"`
	Room         Room       `gorm:"foreignKey:room_id;references:ID"`
	RoomID       uint       `gorm:"column:room_id"`
    User         User       `gorm:"foreignKey:UserID;references:ID"`
    UserID       uint       `gorm:"column:user_id"`
}

func (FeedBack) TableName() string {
	return "FEEDBACK"
}
