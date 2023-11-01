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
	NameFeedBack string    `gorm:"column:name_feed_back;type:varchar(100);default:''"`
	Description  string    `gorm:"column:description;type:text;default:''"`
	Status       string    `gorm:"column:status;type:varchar(24);default:'SEND'"`
	TimeStarted  time.Time `gorm:"column:time_started;default:null"`
	TimeFinish   time.Time `gorm:"column:time_finish;default:null"`
	Room         Room      `gorm:"foreignKey:room_id;references:ID"`
	RoomID       uint      `gorm:"column:room_id"`
	User         User      `gorm:"foreignKey:UserID;references:ID"`
	UserID       uint      `gorm:"column:user_id"`
	Category     Category  `gorm:"foreignKey:category_id;references:ID"`
	CategoryID   uint      `gorm:"column:category_id"`
	Reciever     User      `gorm:"foreignKey:reciever_id;references:ID"`
	RecieverID   uint      `gorm:"column:reciever_id;default:null"`
}

func (FeedBack) TableName() string {
	return "FEEDBACK"
}
