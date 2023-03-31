package model

import "time"

type Topic struct {
	ID            uint64 `gorm:"PrimaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TopicName     string `gorm:"type:varchar(255);not null;" json:"topic_name"`
	UserId        string `gorm:"type:varchar(255);not null" json:"user_id"`
	Content       string `gorm:"type:longtext" json:"content"`
	ParentTopicId string `gorm:"type:varchar(255)" json:"parent_topic_id"`
	Description   string `gorm:"type:longtext" json:"description"`
}
