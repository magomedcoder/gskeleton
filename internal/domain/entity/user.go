package entity

import "strings"

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

func (u *User) IsValidEmail() bool {
	// Простой пример валидации email
	return strings.Contains(u.Email, "@")
}
