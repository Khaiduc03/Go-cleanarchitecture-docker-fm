package entities

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomName string `gorm:"column:room_name;type:varchar(100);default:''"`
	Floor    int    `gorm:"column:floor;type:int;default:1"`
	Building string `gorm:"column:building;type:varchar(100);default:''"`
	Status   int   `gorm:"column:status;default:1"`
}

func (Room) TableName() string {
	return "ROOM"
}
