package entity

import "strings"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func (u *User) IsValidEmail() bool {
	// Простой пример валидации email
	return strings.Contains(u.Email, "@")
}
