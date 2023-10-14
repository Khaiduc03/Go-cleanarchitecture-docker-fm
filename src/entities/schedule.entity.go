package entities

import (
	"gorm.io/gorm"
)

const (
	MORNING       = 1
	MID_MORNING   = 2
	AFTERNOON     = 3
	MID_AFTERNOON = 4
	EVENING       = 5
	MID_EVENING   = 6
	NIGHT         = 7
)

type Schedule struct {
	gorm.Model
	IDRoomID uint   `gorm:"column:id_room"`
	IDUserID uint   `gorm:"column:id_user"`
	IDRoom   Room   `gorm:"foreignKey:IDRoomID;references:ID"`
	IDUser   User   `gorm:"foreignKey:IDUserID;references:ID"`
	Date     int64  `gorm:"column:date;type:bigint;default:0"`
	Shift    uint32 `gorm:"column:shift;type:int;default:1"`
}

func (Schedule) TableName() string {
	return "SCHEDULE"
}
