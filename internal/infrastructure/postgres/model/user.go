package model

import "time"

type User struct {
	Id        int       `gorm:"primaryKey"`
	Username  string    `gorm:"username"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (User) TableName() string {
	return "users"
}
