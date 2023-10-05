package entities

import "gorm.io/gorm"

const (
	ADMIN   = "ADMIN"
	STAFF   = "STAFF"
	TEACHER = "TEACHER"
)

const (
	HCM = "HCM"
	HN  = "HN"
	DN  = "DN"
)

type User struct {
	gorm.Model
	Email       string `gorm:"column:email;type:varchar(100);unique"`
	Dob         int64  `gorm:"column:dob;type:bigint;default:0"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(12);default:''"`
	Url         string `gorm:"column:url;type:varchar(200);default:''"`
	DeviceToken string `gorm:"column:device_token;type:varchar(500);default:''"`
	Role        string `gorm:"column:role;type:varchar(10);default:'USER'"`
	Position    string `gorm:"column:position;type:varchar(2);default:'HCM'"`
	Department  string `gorm:"column:department;type:varchar(2);default:'HCM'"`
}

func (User) TableName() string {
	return "USER"
}
