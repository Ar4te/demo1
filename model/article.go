package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Article struct {
	ID              uint64 `gorm:"PrimaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	FileName        string `gorm:"type:varchar(255);not null;"`
	UserId          string `gorm:"type:varchar(255);not null"`
	FileStream      string `gorm:"type:longtext"`
	ParentArticleId string `gorm:"type:varchar(255)"`
	Description     string `gorm:"type:longtext"`
}
