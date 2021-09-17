package models

import "time"

type User struct {
	ID         uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null"`
	UserID     int64     `gorm:"unique;column:user_id;type:bigint;not null"`
	Username   string    `gorm:"unique;column:username;type:varchar(64);not null"`
	Password   string    `gorm:"column:password;type:varchar(64);not null"`
	Email      string    `gorm:"column:email;type:varchar(64)"`
	Gender     int8      `gorm:"column:gender;type:tinyint;default:0"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP"`
}
