package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	//gorm.Model
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password  string `gorm:"size:255;not null"`
	IsDeleted bool   `gorm:"default:false"`
}
