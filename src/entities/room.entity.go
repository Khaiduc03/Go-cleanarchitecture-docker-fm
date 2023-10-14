package entities

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomName string `gorm:"column:room_name;type:varchar(100);default:''"`
	Floor    int    `gorm:"column:floor;type:int;default:0"`
	Building string `gorm:"column:building;type:varchar(100);default:''"`
	Status   string `gorm:"column:status;type:varchar(100);default:''"`
}

func (Room) TableName() string {
	return "ROOM" // Đổi tên bảng thành "rooms" (thường là dạng số nhiều)
}
