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
	Name        string `gorm:"column:name;type:varchar(100); default:''"`
	Email       string `gorm:"column:email;type:varchar(100);unique"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(12);default:''"`
	Url         string `gorm:"column:url;type:varchar(250);default:''"`
	Role        string `gorm:"column:role;type:varchar(10);default:'TEACHER'"`
	Position    string `gorm:"column:position;type:varchar(50);default:'FPT Polytechnic Hồ Chí Minh'"`
	Department  string `gorm:"column:department;type:varchar(20);default:'Toa T'"`
}

func (User) TableName() string {
	return "USER"
}
