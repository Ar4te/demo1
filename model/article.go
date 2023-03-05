package model

import (
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

type Article struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
	FileName string `gorm:"type:varchar(255);not null;"`
	UserId string `gorm:"type:varchar(255);not null"`
	FileStream string `gorm:"type:longtext"`
}