package entities

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomName string `gorm:"column:room_name; type:varchar(100); default:''"`
}

func (Room) TableName() string {
	return "ROOM"
}
