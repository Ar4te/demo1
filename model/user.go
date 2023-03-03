package model

import (
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
}