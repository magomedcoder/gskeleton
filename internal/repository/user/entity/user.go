package entity

import (
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
