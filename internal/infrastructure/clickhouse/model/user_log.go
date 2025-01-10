package model

import "time"

type UserLog struct {
	UserId    int64     `ch:"user_id"`
	Log       string    `ch:"log"`
	CreatedAt time.Time `ch:"created_at"`
}

func (UserLog) TableName() string {
	return "user_logs"
}
